package model

type Profile struct {
	Name       string
	Gender     string
	Marriage   string
	Age        string
	Height     string
	Weight     string
	Income     string
	Education  string
	Occupation string
	House      string
	Car        string
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
