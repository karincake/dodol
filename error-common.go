package dodol

type CommonError struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail"`
}

func (obj CommonError) Error() string {
	msg := "code: " + obj.Code
	if obj.Message != "" {
		msg += "; message: " + obj.Message
	}
	if obj.Detail != "" {
		msg += "; detail: " + obj.Detail
	}
	return msg
}

type CommonErrors map[string]CommonError

// Get one error by key
func (obj CommonErrors) Error() string {
	output := ""
	for key, err := range obj {
		output += "key: " + key + ", code: " + err.Code
		if err.Message != "" {
			output += "; message: " + err.Message
		}
		if err.Detail != "" {
			output += "; detail: " + err.Detail
		}
		output += "\n"
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
