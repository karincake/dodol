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
	msg := "code: " + obj.Code
	if obj.Message != "" {
		msg += "; message: " + obj.Message
	}
	if obj.ExpectedVal != "" {
		msg += "; expected value: " + obj.ExpectedVal
		msg += "; given value: " + fmt.Sprintf("%v", obj.GivenVal)
	}
	if obj.EmbedSource != "" {
		msg += "; embed source: " + obj.EmbedSource
	}
	return msg
}

type FieldErrors map[string]FieldError

// Get one error by key
func (obj FieldErrors) Error() string {
	output := ""
	for key, err := range obj {
		output += "key: " + key + ", code: " + err.Code
		if err.Message != "" {
			output += "; message: " + err.Message
		}
		if err.ExpectedVal != "" {
			output += "; expected value: " + err.ExpectedVal
			output += "; given value: " + fmt.Sprintf("%v", err.GivenVal)
		}
		if err.EmbedSource != "" {
			output += "; embed source: " + err.EmbedSource
		}
		output += "\n"
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
