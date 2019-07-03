package helpers

type HTTPSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HTTPError struct {
	Error int `json:"error"`
}
