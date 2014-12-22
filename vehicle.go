package swapi

import "fmt"

// A Vehicle is a single transport craft that does not have hyperdrive capability.
type Vehicle struct {
	Name                 string      `json:"name"`
	Model                string      `json:"model"`
	Manufacturer         string      `json:"manufacturer"`
	CostInCredits        string      `json:"cost_in_credits"`
	Length               string      `json:"length"`
	MaxAtmospheringSpeed string      `json:"max_atmosphering_speed"`
	Crew                 string      `json:"crew"`
	Passengers           string      `json:"passengers"`
	CargoCapacity        string      `json:"cargo_capacity"`
	Consumables          string      `json:"consumables"`
	VehicleClass         string      `json:"vehicle_class"`
	PilotURLs            []personURL `json:"pilots"`
	FilmURLs             []filmURL   `json:"films"`
	Created              string      `json:"created"`
	Edited               string      `json:"edited"`
	URL                  string      `json:"url"`
}

type vehicleURL string

// Vehicle retrieves the vehicle with the given id
func (c *Client) Vehicle(id int) (Vehicle, error) {
	req, err := c.NewRequest(fmt.Sprintf("vehicles/%d", id))
	if err != nil {
		return Vehicle{}, err
	}

	var vehicle Vehicle
	if _, err = c.Do(req, &vehicle); err != nil {
		return Vehicle{}, err
	}

	return vehicle, nil
}
