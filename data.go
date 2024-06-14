package dodol

// String indexed of string
type IS map[string]string

// String indexed of IS string indexed of string
type ISIS map[string]IS

// String indexed of interface
type II map[string]interface{}

// Some structured data
type Data struct {
	Meta any `json:"meta,omitempty"`
	Data any `json:"data"`
	Ref  any `json:"ref,omitempty"`
}

type Message struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

type Content struct {
	Content any `json:"content"`
}
