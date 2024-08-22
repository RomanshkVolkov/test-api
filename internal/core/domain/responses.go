package domain

type RequestValidator[T any] struct {
	Type T
}

type APIResponse[T any, Sch any] struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	Data        T      `json:"data" swaggerignore:"true"`
	SchemaError Sch    `json:"schema" swaggerignore:"true"`
	Error       error  `json:"error"`
}

type RequestInfo struct {
	Host      string `json:"host"`
	IP        string `json:"ip"`
	UserAgent string `json:"agent"`
	UserID    uint   `json:"userID"`
}
