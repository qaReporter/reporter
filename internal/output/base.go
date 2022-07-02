package output

type Output struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}
