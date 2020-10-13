package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Weapon string

const (
	baseballBat Weapon = "baseball bat"
	hatchet     Weapon = "hatchet"
	scythe      Weapon = "scythe"
	machete     Weapon = "machete"
)

var heroes = map[Weapon]string{
	baseballBat: "Tex Willer",
	hatchet:     "Kit Carson",
	scythe:      "Jack Tiger",
	machete:     "Kit Willer",
}

type Zombie struct {
	name   string
	weapon Weapon
}

type ZombieChannel chan Zombie

func randomWeapon() Weapon {
	switch rand.Intn(4) {
	case 0:
		return baseballBat
	case 1:
		return hatchet
	case 2:
		return scythe
	case 3:
		return machete
	}
	panic("Unexpected weapon index")
}

func makeZombie(id int) Zombie {
	return Zombie{
		fmt.Sprintf("zombie-%05d", id),
		randomWeapon(),
	}
}

func spawnZombies(zombieCh ZombieChannel, howMany int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= howMany; i++ {
		time.Sleep(150*time.Millisecond + time.Duration(rand.Intn(300))*time.Millisecond)
		zombieCh <- makeZombie(i)
	}
	close(zombieCh)
}

func fightZombie(ok bool, origin string, z Zombie) {
	if ok {
		fmt.Printf("From: %s | Weapon: %s | Hero: %s\n", origin, z.weapon, heroes[z.weapon])
	}
}

func main() {
	northCh := make(ZombieChannel)
	southCh := make(ZombieChannel)
	eastCh := make(ZombieChannel)
	westCh := make(ZombieChannel)

	var wg sync.WaitGroup

	go spawnZombies(northCh, 14, &wg)
	go spawnZombies(southCh, 8, &wg)
	go spawnZombies(eastCh, 10, &wg)
	go spawnZombies(westCh, 12, &wg)

	wg.Add(4)

	go func() {
		for {
			select {
			case z, ok := <-northCh:
				fightZombie(ok, "north", z)
			case z, ok := <-southCh:
				fightZombie(ok, "south", z)
			case z, ok := <-eastCh:
				fightZombie(ok, "east", z)
			case z, ok := <-westCh:
				fightZombie(ok, "west", z)
			}
		}
	}()

	wg.Wait()
}
