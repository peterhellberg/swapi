package swapi

import (
	"context"
	"fmt"
)

// A Species is a type of person or character within the Star Wars Universe.
type Species struct {
	Name            string   `json:"name"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	SkinColors      string   `json:"skin_colors"`
	HairColors      string   `json:"hair_colors"`
	EyeColors       string   `json:"eye_colors"`
	AverageLifespan string   `json:"average_lifespan"`
	Homeworld       string   `json:"homeworld"`
	Language        string   `json:"language"`
	PeopleURLs      []string `json:"people"`
	FilmURLs        []string `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`
	URL             string   `json:"url"`
}

// Species retrieves the species with the given id
func (c *Client) Species(ctx context.Context, id int) (Species, error) {
	req, err := c.newRequest(ctx, fmt.Sprintf("species/%d", id))
	if err != nil {
		return Species{}, err
	}

	var species Species

	if _, err = c.do(req, &species); err != nil {
		return Species{}, err
	}

	return species, nil
}

func (c *Client) AllSpecies(ctx context.Context) ([]Species, error) {
	var species []Species

	req, err := c.newRequest(ctx, "species/")
	if err != nil {
		return nil, err
	}

	for {
		var list List[Species]

		if _, err = c.do(req, &list); err != nil {
			return nil, err
		}

		species = append(species, list.Results...)

		if list.Next == nil {
			break
		}

		req, err = c.getRequest(ctx, *list.Next)
		if err != nil {
			return nil, err
		}
	}

	return species, nil
}
