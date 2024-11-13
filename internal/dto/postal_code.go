package dto

type PostalCodeRequest struct {
	PostalCode []int `json:"postalCodes"`
}

type PostalCodeResponse struct {
	PostalCode  int    `json:"postalCode"`
	AreaCode    string `json:"areaCode"`
	Province    string `json:"province"`
	City        string `json:"city"`
	SubDistrict string `json:"subDistrict"`
	Village     string `json:"village"`
}
