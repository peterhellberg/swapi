package swapi

import "fmt"

// A Planet is a large mass, planet or planetoid in the Star Wars Universe, at the time of 0 ABY.
type Planet struct {
	Name           string        `json:"name"`
	RotationPeriod string        `json:"rotation_period"`
	OrbitalPeriod  string        `json:"orbital_period"`
	Diameter       string        `json:"diameter"`
	Climate        string        `json:"climate"`
	Gravity        string        `json:"gravity"`
	Terrain        string        `json:"terrain"`
	SurfaceWater   string        `json:"surface_water"`
	Population     string        `json:"population"`
	ResidentURLs   []residentURL `json:"residents"`
	FilmURLs       []filmURL     `json:"films"`
	Created        string        `json:"created"`
	Edited         string        `json:"edited"`
	URL            string        `json:"url"`
}

type planetURL string
type residentURL string

// Planet retrieves the planet with the given id
func (c *Client) Planet(id int) (Planet, error) {
	req, err := c.NewRequest(fmt.Sprintf("planets/%d", id))
	if err != nil {
		return Planet{}, err
	}

	var planet Planet
	if _, err = c.Do(req, &planet); err != nil {
		return Planet{}, err
	}

	return planet, nil
}
