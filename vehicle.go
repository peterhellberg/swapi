package swapi

import (
	"context"
	"fmt"
)

// A Vehicle is a single transport craft that does not have hyperdrive capability.
type Vehicle struct {
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
	VehicleClass         string   `json:"vehicle_class"`
	PilotURLs            []string `json:"pilots"`
	FilmURLs             []string `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  string   `json:"url"`
}

// Vehicle retrieves the vehicle with the given id
func (c *Client) Vehicle(ctx context.Context, id int) (Vehicle, error) {
	req, err := c.newRequest(ctx, fmt.Sprintf("vehicles/%d", id))
	if err != nil {
		return Vehicle{}, err
	}

	var vehicle Vehicle

	if _, err = c.do(req, &vehicle); err != nil {
		return Vehicle{}, err
	}

	return vehicle, nil
}

func (c *Client) AllVehicles(ctx context.Context) ([]Vehicle, error) {
	var vehicles []Vehicle

	req, err := c.newRequest(ctx, "vehicles/")
	if err != nil {
		return nil, err
	}

	for {
		var list List[Vehicle]

		if _, err = c.do(req, &list); err != nil {
			return nil, err
		}

		vehicles = append(vehicles, list.Results...)

		if list.Next == nil {
			break
		}

		req, err = c.getRequest(ctx, *list.Next)
		if err != nil {
			return nil, err
		}
	}

	return vehicles, nil
}
