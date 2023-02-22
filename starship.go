package swapi

import (
	"context"
	"fmt"
)

// A Starship is a single transport craft that has hyperdrive capability.
type Starship struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	HyperdriveRating     string   `json:"hyperdrive_rating"`
	MGLT                 string   `json:"MGLT"`
	StarshipClass        string   `json:"starship_class"`
	PilotURLs            []string `json:"pilots"`
	FilmURLs             []string `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  string   `json:"url"`
}

// Starship retrieves the starship with the given id
func (c *Client) Starship(ctx context.Context, id int) (Starship, error) {
	req, err := c.newRequest(ctx, fmt.Sprintf("starships/%d", id))
	if err != nil {
		return Starship{}, err
	}

	var starship Starship

	if _, err = c.do(req, &starship); err != nil {
		return Starship{}, err
	}

	return starship, nil
}

func (c *Client) AllStarships(ctx context.Context) ([]Starship, error) {
	var starships []Starship

	req, err := c.newRequest(ctx, "starships/")
	if err != nil {
		return nil, err
	}

	for {
		var list List[Starship]

		if _, err = c.do(req, &list); err != nil {
			return nil, err
		}

		starships = append(starships, list.Results...)

		if list.Next == nil {
			break
		}

		req, err = c.getRequest(ctx, *list.Next)
		if err != nil {
			return nil, err
		}
	}

	return starships, nil
}
