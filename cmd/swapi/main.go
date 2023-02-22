package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/peterhellberg/swapi"
)

func main() {
	c := swapi.DefaultClient

	if len(os.Args) < 2 {
		usage()
		return
	}

	command := os.Args[1]

	if len(os.Args) > 2 {
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return
		}

		switch command {
		case "film":
			dump(c.Film(id))
		case "person":
			dump(c.Person(id))
		case "planet":
			dump(c.Planet(id))
		case "species":
			dump(c.Species(id))
		case "starship":
			dump(c.Starship(id))
		case "vehicle":
			dump(c.Vehicle(id))
		}
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		switch command {
		case "people":
			dump(c.AllPeople(ctx))
		case "planets":
			dump(c.AllPlanets(ctx))
		case "films":
			dump(c.AllFilms(ctx))
		case "species":
			dump(c.AllSpecies(ctx))
		case "vehicles":
			dump(c.AllVehicles(ctx))
		case "starships":
			dump(c.AllStarships(ctx))
		}
	}
}

func usage() {
	fmt.Println(strings.Join([]string{
		"Commands:",
		"film     [id]",
		"person   [id]",
		"planet   [id]",
		"species  [id]",
		"starship [id]",
		"vehicle  [id]",
		"films",
		"people",
		"planets",
		"films",
		"species",
		"vehicles",
		"starships",
	}, "\n\t"))
}

func dump(data interface{}, err error) {
	if err != nil {
		return
	}

	if j, err := json.MarshalIndent(data, "", "  "); err == nil {
		fmt.Printf("%s\n", j)
	}
}
