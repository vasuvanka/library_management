package models

type GeneralResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

type CreateUser struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname,omitempty"`
	Username string `json:"username"`
}

type CreateBook struct {
	Author string `json:"author"`
	Name string `json:"name"`
	Copies int `json:"copies"`
}

type ResponseError struct {
    status  int
    message string
}

func (e *ResponseError) Error() string  {
	return e.message
}