package handler

import "sport/sportclub"

type getAllComplexResponse struct {
	Data []sportclub.ComplexJSON `json:"data"`
}

type statusResponse struct {
	Status string `json:"status"`
}
