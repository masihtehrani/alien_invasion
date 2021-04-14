package usecases

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/masihtehrani/alien_invasion/internal/entities/structs"
)

func choseAlienName(_ context.Context, i int) string {
	names := [10]string{
		"Liam",
		"Olivia",
		"Noah",
		"Emma",
		"Oliver",
		"Ava",
		"William",
		"Sophia",
		"Elijah",
		"Isabella",
	}

	m := i % len(names)
	name := names[m]
	d := i / len(names)
	name += fmt.Sprintf("_%d", d)

	return name
}

func (u *UseCases) createAliens(ctx context.Context) {
	for i := 0; i < u.config.alienCount; i++ {
		newAlien := structs.Alien{
			Name: choseAlienName(ctx, i),
		}
		u.aliens[newAlien.Name] = newAlien
	}
}

func (u *UseCases) assignAliensToCities(ctx context.Context) {
	for key := range u.aliens {
		if len(u.worlds) == 0 {
			return
		}

		alien := u.aliens[key]

		city := u.randomCity(ctx)

		alien.City = city
		u.aliens[key] = alien
		u.worlds[city.Name].Aliens[alien.Name] = alien

		u.checkDestroyedCity(ctx, city)
	}
}

func (u *UseCases) randomCity(_ context.Context) *structs.City {
	var i int

	worldCount := len(u.worlds)
	if worldCount > 0 {
		rand.Seed(time.Now().UnixNano())
		i = rand.Intn(len(u.worlds))
	}

	for k := range u.worlds {
		if i == 0 {
			return u.worlds[k]
		}

		i--
	}

	return nil
}

func (u *UseCases) checkDestroyedCity(ctx context.Context, city *structs.City) {
	if len(city.Aliens) == AlienFightCondition {
		aliensName := make([]string, 2)

		var i int

		for k := range city.Aliens {
			aliensName[i] = k
			delete(u.aliens, k)

			i++
		}

		u.logger.Info.Printf(" city %s has been destroyed by alien %s and alien %s!\n", city.Name, aliensName[0],
			aliensName[1])

		for i := range city.Links {
			delete(city.Links[i].Links, oppositeDirections(ctx, i))
		}

		delete(u.worlds, city.Name)
	}
}
