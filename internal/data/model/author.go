package model

type Author struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Books     []Book `json:"books"`
}
