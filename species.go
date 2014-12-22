package swapi

import "fmt"

type Species struct {
	Name            string      `json:"name"`
	Classification  string      `json:"classification"`
	Designation     string      `json:"designation"`
	AverageHeight   string      `json:"average_height"`
	SkinColors      string      `json:"skin_colors"`
	HairColors      string      `json:"hair_colors"`
	EyeColors       string      `json:"eye_colors"`
	AverageLifespan string      `json:"average_lifespan"`
	Homeworld       string      `json:"homeworld"`
	Language        string      `json:"language"`
	PeopleURLs      []personURL `json:"people"`
	FilmURLs        []filmURL   `json:"films"`
	Created         string      `json:"created"`
	Edited          string      `json:"edited"`
	URL             string      `json:"url"`
}

type speciesURL string

func (c *Client) Species(id int) (Species, error) {
	req, err := c.NewRequest(fmt.Sprintf("species/%d", id))
	if err != nil {
		return Species{}, err
	}

	var species Species
	if _, err = c.Do(req, &species); err != nil {
		return Species{}, err
	}

	return species, nil
}
