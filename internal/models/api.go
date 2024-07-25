package models

// HomeV1Response - api v1 home route response
type HomeV1Response struct {
	Version     string `json:"version"`
	Description string `json:"description"`
}

// JSONResponse - default response
type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// RunV1Request - api v1 run route request
type RunV1Request struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

// CheckV1Response - api v1 check route response
type CheckV1Response struct {
	Status string `json:"status"`
	Result string `json:"result"`
}
