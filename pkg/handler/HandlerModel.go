package handler

import "sport/sportclubs"

type getAllComplexResponse struct {
	Data []sportclubs.ComplexJSON `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}
