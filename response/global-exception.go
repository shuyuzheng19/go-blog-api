package response

type GlobalException struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewGlobalException(code int, message string) GlobalException {
	return GlobalException{
		Code:    code,
		Message: message,
	}
}
