package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Weapon string
type Hero string

const (
	baseballBat Weapon = "baseball bat"
	hatchet     Weapon = "hatchet"
	scythe      Weapon = "scythe"
	machete     Weapon = "machete"
)

const (
	tex     Hero = "Tex Willer"
	kcarson Hero = "Kit Carson"
	jack    Hero = "Jack Tiger"
	kwiller Hero = "Kit Willer"
)

const (
	timeout = 5 * time.Second
)

var heroes = map[Weapon]Hero{
	baseballBat: tex,
	hatchet:     kcarson,
	scythe:      jack,
	machete:     kwiller,
}

type Zombie struct {
	name   string
	weapon Weapon
}

type ZombieChannel chan Zombie
type HeroChannel chan Zombie

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

func makeZombie(id int, prefix string) Zombie {
	return Zombie{
		fmt.Sprintf("%s-%05d", prefix, id),
		randomWeapon(),
	}
}

func delay(baseMs int, additionalMs int) {
	time.Sleep(time.Duration(baseMs+rand.Intn(additionalMs)) * time.Millisecond)
}

func spawnZombies(howMany int, prefix string, zombieCh ZombieChannel) {
	for i := 1; i <= howMany; i++ {
		delay(150, 1000)
		zombieCh <- makeZombie(i, prefix)
	}
	close(zombieCh)
}

func fightZombie(ok bool, z Zombie, h Hero) {
	if ok {
		fmt.Printf("[FIGHT] %s confronts %s\n", h, z.name)
		delay(250, 1500)
		fmt.Printf(" [KILL] %s defeats %s\n", h, z.name)
	}
}

func main() {
	northCh := make(ZombieChannel)
	southCh := make(ZombieChannel)
	eastCh := make(ZombieChannel)
	westCh := make(ZombieChannel)

	texCh := make(HeroChannel, 5)
	kcarsonCh := make(HeroChannel, 5)
	jackCh := make(HeroChannel, 5)
	kwillerCh := make(HeroChannel, 5)

	terminationCh := make(chan bool)

	var t0 = time.Now()

	go spawnZombies(20, "N", northCh)
	go spawnZombies(18, "S", southCh)
	go spawnZombies(12, "E", eastCh)
	go spawnZombies(22, "W", westCh)

	go func() {
		lastSpawn := time.Now()
		for {
			var z Zombie
			var ok bool
			select {
			case z, ok = <-northCh:
			case z, ok = <-southCh:
			case z, ok = <-eastCh:
			case z, ok = <-westCh:
			}
			if ok {
				fmt.Printf("[SPAWN] %s with a %s\n", z.name, z.weapon)
				lastSpawn = time.Now()
				switch heroes[z.weapon] {
				case tex:
					texCh <- z
				case kcarson:
					kcarsonCh <- z
				case jack:
					jackCh <- z
				case kwiller:
					kwillerCh <- z
				}
			}
			if time.Since(lastSpawn) > timeout {
				close(kwillerCh)
				close(jackCh)
				close(kcarsonCh)
				close(texCh)
				terminationCh <- true
				return
			}
		}
	}()

	go func() {
		for {
			z, ok := <-texCh
			fightZombie(ok, z, tex)
		}
	}()

	go func() {
		for {
			z, ok := <-kcarsonCh
			fightZombie(ok, z, kcarson)
		}
	}()

	go func() {
		for {
			z, ok := <-jackCh
			fightZombie(ok, z, jack)
		}
	}()

	go func() {
		for {
			z, ok := <-kwillerCh
			fightZombie(ok, z, kwiller)
		}
	}()

	<-terminationCh

	tn := time.Since(t0)
	fmt.Printf("Outbreak lasted for %dms\n", tn.Milliseconds())
}
