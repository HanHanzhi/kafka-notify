package models

type User struct {
	ID   int    `json:"id"`   //specifies that ID field in struct should be marshaled(convert to json) with the key "id"
	Name string `json:"name"` //we do this when we need JSON keys to differ from the struct field names
}

type Notification struct {
	From    User   `json:"from"`
	To      User   `json:"to"`
	Message string `json:"message"`
}
