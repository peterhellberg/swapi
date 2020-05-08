package swapi

import "fmt"

// A Person is an individual person or character within the Star Wars universe.
type Person struct {
	Name         string   `json:"name"`
	Height       string   `json:"height"`
	Mass         string   `json:"mass"`
	HairColor    string   `json:"hair_color"`
	SkinColor    string   `json:"skin_color"`
	EyeColor     string   `json:"eye_color"`
	BirthYear    string   `json:"birth_year"`
	Gender       string   `json:"gender"`
	Homeworld    string   `json:"homeworld"`
	FilmURLs     []string `json:"films"`
	SpeciesURLs  []string `json:"species"`
	VehicleURLs  []string `json:"vehicles"`
	StarshipURLs []string `json:"starships"`
	Created      string   `json:"created"`
	Edited       string   `json:"edited"`
	URL          string   `json:"url"`
}

// Person retrieves the person with the given id
func (c *Client) Person(id int) (Person, error) {
	req, err := c.newRequest(fmt.Sprintf("people/%d", id))
	if err != nil {
		return Person{}, err
	}

	var person Person

	if _, err = c.do(req, &person); err != nil {
		return Person{}, err
	}

	return person, nil
}
