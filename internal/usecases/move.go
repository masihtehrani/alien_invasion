package usecases

import (
	"context"
	"math/rand"
	"time"

	"github.com/masihtehrani/alien_invasion/internal/entities/structs"
)

func (u *UseCases) MoveAliens(ctx context.Context) {
	for _, alien := range u.aliens {
		if len(u.worlds) == 0 {
			break
		}

		if len(alien.City.Links) == 0 {
			continue
		}

		u.move(&alien)
		u.checkDestroyedCity(ctx, alien.City)
	}
}

func (u *UseCases) move(alien *structs.Alien) {
	city := u.randomLink(alien.City)
	alien.City = city
	u.aliens[alien.Name] = *alien
	u.worlds[city.Name].Aliens[alien.Name] = *alien
	delete(u.worlds[alien.City.Name].Aliens, alien.Name)
	delete(u.aliens, alien.Name)
}

func (u *UseCases) randomLink(city *structs.City) *structs.City {
	var i int

	linkCount := len(city.Links)
	if linkCount > 0 {
		rand.Seed(time.Now().UnixNano())
		i = rand.Intn(linkCount)
	}

	for k := range city.Links {
		if i == 0 {
			return city.Links[k]
		}

		i--
	}

	return nil
}
