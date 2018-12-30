package skyscanner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func prettyPrint(apiResponse Response) {
	fmt.Printf("%+v\n", apiResponse.Quotes)
}

func parseResponse(response *http.Response) Response {
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	apiResponse := Response{}
	jsonErr := json.Unmarshal(body, &apiResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return apiResponse
}

func getRequest(url string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("X-Mashape-Key", *apiConfig.MashapeKey)

	if err != nil {
		log.Fatal(err)
	}
	return req
}

func getClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 2,
	}
}

func formatURL(path string) string {
	baseURL := *apiConfig.BaseURL

	return fmt.Sprintf("%v%v", baseURL, path)
}

func get(url string) *http.Response {
	res, getErr := getClient().Do(getRequest(url))

	if getErr != nil {
		log.Fatal(getErr)
	}

	return res

}

/*
BrowseQuotes stub
*/
func BrowseQuotes(parameters BrowseParameters) Response {
	browseQuotes := formatURL(fmt.Sprintf("browsequotes/v1.0/%v/%v/%v/%v/%v/%v/%v",
		parameters.Country,
		parameters.Currency,
		parameters.Locale,
		parameters.OriginPlace,
		parameters.DestinationPlace,
		parameters.OutboundDate,
		parameters.InboundDate,
	))
	res := get(browseQuotes)
	return parseResponse(res)
}

// ProcessDestination does asynchronous api requests to the SkyScanner API
func ProcessDestination(params *BrowseParameters) (*QuoteSummary, error) {
	SkyscannerQuotes := BrowseQuotes(*params)
	quote, err := SkyscannerQuotes.LowestPrice()
	if err != nil {
		log.Printf("%v\n\n", err)
		return nil, err
	}
	return quote, nil
}

// ProcessDestinationAsync asynchronously processes quote requests
func ProcessDestinationAsync(params *BrowseParameters, out chan<- *QuoteSummary) {
	q, err := ProcessDestination(params)
	if err != nil {
		out <- nil
	}
	out <- q
}

// BatchDestinations processes multiple destinations
func BatchDestinations(params *BatchBrowseParameters) []*QuoteSummary {
	var newQuotes []*QuoteSummary
	quoteChannels := make(chan *QuoteSummary)
	for _, origin := range params.Origins {
		for _, destination := range params.Destinations {
			fmt.Println("origin", origin, "destination", destination)
			p := BrowseParameters{
				BaseParameters:   params.BaseParameters,
				DestinationPlace: destination,
				OriginPlace:      origin,
			}
			go ProcessDestinationAsync(&p, quoteChannels)
		}
	}
	for range params.Destinations {
		q := <-quoteChannels
		fmt.Println(q)
		if q == nil {
			continue
		}
		newQuotes = append(newQuotes, q)
	}
	return newQuotes
}
