package domain

type BrigateResponse[T any] struct {
	StatusCode      int         `json:"statusCode"`
	ErrorCode       string      `json:"errorCode"`
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	Errors          interface{} `json:"errors"`
	Data            []T         `json:"data"`
}
