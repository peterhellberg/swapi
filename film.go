package swapi

import "fmt"

// A Film is an single film.
type Film struct {
	Title         string         `json:"title"`
	EpisodeID     int            `json:"episode_id"`
	OpeningCrawl  string         `json:"opening_crawl"`
	Director      string         `json:"director"`
	Producer      string         `json:"producer"`
	CharacterURLs []characterURL `json:"characters"`
	PlanetURLs    []planetURL    `json:"planets"`
	StarshipURLs  []starshipURL  `json:"starships"`
	VehicleURLs   []vehicleURL   `json:"vehicles"`
	SpeciesURLs   []speciesURL   `json:"species"`
	Created       string         `json:"created"`
	Edited        string         `json:"edited"`
	URL           string         `json:"url"`
}

type filmURL string
type characterURL string

// Film retrieves the film with the given id
func (c *Client) Film(id int) (Film, error) {
	req, err := c.NewRequest(fmt.Sprintf("films/%d", id))
	if err != nil {
		return Film{}, err
	}

	var film Film
	if _, err = c.Do(req, &film); err != nil {
		return Film{}, err
	}

	return film, nil
}
