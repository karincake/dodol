package dodol

import "fmt"

type CommonError struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

func (obj CommonError) Error() string {
	return fmt.Sprintf("code: %v; message: %v", obj.Code, obj.Message)
}

type CommonErrors map[string]CommonError

// Get one error by key
func (obj CommonErrors) Error() string {
	output := ""
	for key, err := range obj {
		output += fmt.Sprintf("key: %v, code: %v; message: %v\n", key, err.Code, err.Message)
	}
	return output
}

// Get the first error by key
func (obj CommonErrors) GetFirst() (string, error) {
	for key, err := range obj {
		return key, err
	}
	return "", nil
}

// Check if a key exists
func (obj CommonErrors) KeyExists(key string) bool {
	if _, ok := obj[key]; ok {
		return true
	}
	return false
}

// Import list from other CommonErrors
func (obj CommonErrors) Import(src CommonErrors) {
	for idx, val := range src {
		obj[idx] = val
	}
}
