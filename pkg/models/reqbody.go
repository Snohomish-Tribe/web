package models

type Reqbody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Question string `json:"questions"`
	Message  string `json:"message"`
}
