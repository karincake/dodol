package dodol

import "fmt"

type FieldError struct {
	Source      string `json:"source,omitempty"`
	Code        string `json:"code"`
	Message     string `json:"message,omitempty"`
	ExpectedVal string `json:"expectedVal,omitempty"`
	GivenVal    any    `json:"givenVal,omitempty"`
	EmbedSource string `json:"embedSource,omitempty"`
}

func (obj FieldError) Error() string {
	return fmt.Sprintf("code: %v; message: %v; expected value: %v; given value: %v", obj.Code, obj.Message, obj.ExpectedVal, obj.GivenVal)
}

type FieldErrors map[string]FieldError

// Get one error by key
func (obj FieldErrors) Error() string {
	output := ""
	for key, err := range obj {
		output += fmt.Sprintf("key: %v, code: %v; message: %v; expected value: %v; given value: %v \n", key, err.Code, err.Message, err.ExpectedVal, err.GivenVal)
	}
	return output
}

// Get the first error by key
func (obj FieldErrors) GetFirst() (string, error) {
	for key, err := range obj {
		return key, err
	}
	return "", nil
}

// Check if a key exists
func (obj FieldErrors) KeyExists(key string) bool {
	if _, ok := obj[key]; ok {
		return true
	}
	return false
}

// Import list from other FieldErrors
func (obj FieldErrors) Import(src FieldErrors) {
	for idx, val := range src {
		obj[idx] = val
	}
}
