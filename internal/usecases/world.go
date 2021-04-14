package usecases

import (
	"context"
	"strings"

	"github.com/masihtehrani/alien_invasion/internal/entities/structs"
)

func (u *UseCases) createWorlds(ctx context.Context, textWorlds []string) {
	for i := range textWorlds {
		textCities := strings.Split(textWorlds[i], " ")
		firstCity := strings.ToLower(textCities[0])

		u.worlds[firstCity] = &structs.City{
			Name:   firstCity,
			Links:  make(map[string]*structs.City),
			Aliens: make(map[string]structs.Alien),
		}

		for i := 1; i < len(textCities); i++ {
			linkData := strings.Split(textCities[i], "=")
			direction := strings.ToLower(linkData[0])
			cityName := strings.ToLower(linkData[1])

			u.worlds[cityName] = &structs.City{
				Name:   cityName,
				Links:  make(map[string]*structs.City),
				Aliens: make(map[string]structs.Alien),
			}

			u.worlds[cityName].Links[oppositeDirections(ctx, direction)] = u.worlds[firstCity]

			u.worlds[firstCity].Links[direction] = u.worlds[cityName]
		}
	}
}

func oppositeDirections(_ context.Context, direction string) string {
	compass := map[string]string{
		"north": "south",
		"south": "north",
		"west":  "east",
		"east":  "west",
	}

	return compass[direction]
}

func (u *UseCases) printWorld() []string {
	result := make([]string, len(u.worlds))

	var i int

	for _, v := range u.worlds {
		var links []string

		for k, v := range v.Links {
			links = append(links, k+"="+v.Name)
		}

		result[i] = v.Name + " " + strings.Join(links, " ")

		i++
	}

	return result
}
