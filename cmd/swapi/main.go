package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/peterhellberg/swapi"
)

func main() {
	c := swapi.DefaultClient

	if len(os.Args) < 2 {
		usage()
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return
	}

	switch os.Args[1] {
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
