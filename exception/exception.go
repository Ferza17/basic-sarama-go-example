package exception

// Code ...
type Code string

// Exception ...
type Exception struct {
	ErrorCode     string `json:"error_code"`
	StatusCode    int    `json:"status_code"`
	Message       string `json:"message"`
	CustomMessage string `json:"custom_message"`
	Error         error  `json:"-"`
}

var except = make(map[Code]Exception)

// GetException ...
func GetException(c Code) *Exception {
	e := except[c]
	if e.CustomMessage != "" {
		e.Message = e.CustomMessage
	}
	return &e
}

// GetExceptionCustomMessage ...
func GetExceptionCustomMessage(c Code, msg string) *Exception {
	e := except[c]
	e.Message = msg
	if e.CustomMessage != "" {
		e.Message = e.CustomMessage
	}
	return &e
}
