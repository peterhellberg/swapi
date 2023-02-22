package swapi

import (
	"context"
	"fmt"
)

// A Film is an single film.
type Film struct {
	Title         string   `json:"title"`
	EpisodeID     int      `json:"episode_id"`
	OpeningCrawl  string   `json:"opening_crawl"`
	Director      string   `json:"director"`
	Producer      string   `json:"producer"`
	CharacterURLs []string `json:"characters"`
	PlanetURLs    []string `json:"planets"`
	StarshipURLs  []string `json:"starships"`
	VehicleURLs   []string `json:"vehicles"`
	SpeciesURLs   []string `json:"species"`
	Created       string   `json:"created"`
	Edited        string   `json:"edited"`
	URL           string   `json:"url"`
}

// Film retrieves the film with the given id
func (c *Client) Film(ctx context.Context, id int) (Film, error) {
	req, err := c.newRequest(ctx, fmt.Sprintf("films/%d", id))
	if err != nil {
		return Film{}, err
	}

	var film Film

	if _, err = c.do(req, &film); err != nil {
		return Film{}, err
	}

	return film, nil
}

func (c *Client) AllFilms(ctx context.Context) ([]Film, error) {
	var films []Film

	req, err := c.newRequest(ctx, "films/")
	if err != nil {
		return nil, err
	}

	for {
		var list List[Film]

		if _, err = c.do(req, &list); err != nil {
			return nil, err
		}

		films = append(films, list.Results...)

		if list.Next == nil {
			break
		}

		req, err = c.getRequest(ctx, *list.Next)
		if err != nil {
			return nil, err
		}
	}

	return films, nil
}
