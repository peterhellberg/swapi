package swapi

import "fmt"

type Film struct {
	Title         string         `json:"title"`
	EpisodeId     int            `json:"episode_id"`
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
type planetURL string
type characterURL string

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
