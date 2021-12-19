package responses

type ErrorResponse struct {
	Errors map[string][]string `structs:"errors"`
}
