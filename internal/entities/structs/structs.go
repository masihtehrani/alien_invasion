package structs

type City struct {
	Name   string
	Links  map[string]*City
	Aliens map[string]Alien
}

type Alien struct {
	Name string
	City *City
}
