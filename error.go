package utube

import "fmt"

type ResponseError struct {
	Method  string
	URL     string
	Content *Error `json:"error"`
}

func (this *ResponseError) Error() string {
	return fmt.Sprintf("[%s-%s] %s", this.Method, this.URL, this.Content.Message)
}

type Error struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Errors  []*ErrorInfo   `json:"errors"`
	Status  string         `json:"status"`
	Details []*ErrorDetail `json:"details,omitempty"`
}

type ErrorInfo struct {
	Domain       string `json:"domain"`
	Reason       string `json:"reason"`
	Message      string `json:"message"`
	LocationType string `json:"locationType"`
	Location     string `json:"location"`
}

type ErrorDetail struct {
	Type     string         `json:"@type"`
	Reason   string         `json:"reason"`
	Domain   string         `json:"domain"`
	Metadata *ErrorMetadata `json:"metadata,omitempty"`
}

type ErrorMetadata struct {
	Service  string `json:"service"`
	Consumer string `json:"consumer"`
}
