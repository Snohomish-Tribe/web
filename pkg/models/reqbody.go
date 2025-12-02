package models

type Reqbody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Question string `json:"questions"`
	Message  string `json:"message"`
}

type EventsPageData struct {
	Data []struct {
		Day        string `json:"Day"`
		Date       string `json:"Date"`
		Month_Year string `json:"Month_Year"`
		Time       string `json:"Time"`
		Event_Name string `json:"Event_Name"`
		Address    string `json:"Address"`
		Details    string `json:"Details"`
	} `json:"data"`
}
