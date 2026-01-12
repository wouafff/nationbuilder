package nationbuilder

// The address resource represents an address within many nationbuilder types such as people
type Address struct {
	FirstLine   string `json:"address1,omitempty"`
	SecondLine  string `json:"address2,omitempty"`
	ThirdLine   string `json:"address3,omitempty"`
	HouseNumber string `json:"street_number,omitempty"`
	Street	    string `json:"street_name,omitempty"`
	StreetPrefix string `json:"street_prefix,omitempty"`
	StreetType   string `json:"street_type,omitempty"`
	StreetSuffix string `json:"street_suffix,omitempty"`
	Addition    string `json:"addition,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	County      string `json:"county,omitempty"`
	ZIPCode     string `json:"zip,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Latitude    string `json:"lat,omitempty"`
	Longtitude  string `json:"lng,omitempty"`
}
