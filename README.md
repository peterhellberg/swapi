A [SWAPI](http://swapi.co/) client written in Go
================================================

[![Build Status](https://travis-ci.org/peterhellberg/swapi.svg?branch=master)](https://travis-ci.org/peterhellberg/swapi)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/swapi)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/swapi#license-mit)

## Installation

```bash
go get -u github.com/peterhellberg/swapi
```

## Examples

### atst.go

```go
package main

import (
	"fmt"

	"github.com/peterhellberg/swapi"
)

func main() {
	c := swapi.DefaultClient

	if atst, err := c.Vehicle(19); err == nil {
		fmt.Println("name: ", atst.Name)
		fmt.Println("model:", atst.Model)
	}
}
```

## Command line tool

### Installation

```bash
go get github.com/peterhellberg/swapi/cmd/swapi
```

### Usage

```bash
Commands:
  film     [id]
  person   [id]
  planet   [id]
  species  [id]
  starship [id]
  vehicle  [id]
```

```json
$ swapi planet 1
{
  "name": "Tatooine",
  "rotation_period": "23",
  "orbital_period": "304",
  "diameter": "10465",
  "climate": "arid",
  "gravity": "1 standard",
  "terrain": "desert",
  "surface_water": "1",
  "population": "200000",
  "residents": [
    "http://swapi.co/api/people/1/",
    "http://swapi.co/api/people/2/",
    "http://swapi.co/api/people/4/",
    "http://swapi.co/api/people/6/",
    "http://swapi.co/api/people/7/",
    "http://swapi.co/api/people/8/",
    "http://swapi.co/api/people/9/",
    "http://swapi.co/api/people/11/",
    "http://swapi.co/api/people/43/",
    "http://swapi.co/api/people/62/"
  ],
  "films": [
    "http://swapi.co/api/films/1/",
    "http://swapi.co/api/films/3/",
    "http://swapi.co/api/films/4/",
    "http://swapi.co/api/films/5/",
    "http://swapi.co/api/films/6/"
  ],
  "created": "2014-12-09T13:50:49.641000Z",
  "edited": "2014-12-21T20:48:04.175778Z",
  "url": "http://swapi.co/api/planets/1/"
}
```

## License (MIT)

*Copyright (C) 2014 Peter Hellberg*

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the "Software"),
> to deal in the Software without restriction, including without limitation
> the rights to use, copy, modify, merge, publish, distribute, sublicense,
> and/or sell copies of the Software, and to permit persons to whom the
> Software is furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included
> in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
> OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
> IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
> DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
> TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
> OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
