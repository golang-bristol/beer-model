package model

import "time"

type Review struct {
	ID        int       `json:"id"`
	BeerID    int       `json:"beer_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Score     int       `json:"score"`
	Text      string    `json:"text"`
	Created   time.Time `json:"created"`
}

type Beer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Brewery   string    `json:"brewery"`
	Abv       int       `json:"abv"`
	ShortDesc string    `json:"short_description"`
	Created   time.Time `json:"created"`
}
