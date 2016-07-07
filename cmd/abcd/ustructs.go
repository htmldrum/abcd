package main

type Feed struct {
	Id uint
	Name string
	Description string
	URI string
	Subjects []string
	Networks []string
	Last_contact_datetime string // changing
	Articles []Article
}

type Article struct {
	Name, Url string
}

type Config struct {
	faves []string
	feeds []string
}
