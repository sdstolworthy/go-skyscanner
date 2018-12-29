package skyscanner

// BaseParameters generic request type for skyscanner api
type BaseParameters struct {
	Country     string `json:"country"`
	Currency    string `json:"currency"`
	Locale      string `json:"locale"`
	OutbandDate string `json:"outboundDate"`
	Adults      int    `json:"adults"`
	InboundDate string `json:"inboundDate"`
}

// BrowseParameters represent the parameters needed to access the Browse Quotes SkyScanner API
type BrowseParameters struct {
	BaseParameters
	OriginPlace      string `json:"originPlace"`
	DestinationPlace string `json:"destinationPlace"`
}

// BatchBrowseParameters allows batch requests
type BatchBrowseParameters struct {
	BaseParameters
	Origins      []string `json:"originPlaces"`
	Destinations []string `json:"destinationPlaces"`
}
