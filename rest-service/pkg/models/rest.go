package models

// RestErr is REST API error
type RestErr struct {
	Code   int      `json:"code,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

// PeopleFilter for filtering list of people using query string
type PeopleFilter struct {
	FirstName   string
	LastName    string
	PhoneNumber string
}
