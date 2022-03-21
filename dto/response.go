package dto

type HeadResponse struct {
	Url           string `json:"Url"`
	StatusCode    int    `json:"Status-Code"`
	ContentLength int64  `json:"Content-Length"`
}
