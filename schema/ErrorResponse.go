package schema

type ErrorResponse struct {
	Failed string `json:"failed"`
	Tag    string `json:"tag"`
	Value  string `json:"value"`
}
