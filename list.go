package swapi

type Entity interface {
	Person | Planet | Species | Starship | Vehicle
}

type List[E Entity] struct {
	Count   int
	Next    *string
	Results []E
}
