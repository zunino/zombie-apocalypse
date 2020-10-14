# Zombie Apocalypse

This program was inspired by a [zombie apocalypse problem statement created by Johni](https://gist.github.com/johnidm/3e98e4ff7565ce985a41).
His original intent was to use the scenario as the basis for playing
with golang's concurrent concepts and primitives.

In Johni's conception, zombies use *1 of 4 different types of weapons*
and there are *4 heroes* who are supposed to defeat the invading
zombies. Each hero specializes in one kind of zombie weapon, meaning
they will only fight those carrying the appropriate weapons. 

### Zombie weapons vs. heroes

| Zombie weapon    | Hero       |
|------------------|------------|
| Baseball bat     | Tex Willer |
| Hatchet          | Kit Carson |
| Scythe           | Jack Tiger |
| Machete          | Kit Willer |

## Adaptations

In the original design, channels were used to represent each type of
zombie and their respective weapon. In other words, if a
hatchet-carrying zombie was spawned, it would be sent to the hatchet
channel. The receiving end would then assign the zombie to the
corresponding hero, since it would know its weapon type.

In the present design, there are 2 different types of channels:

  * `ZombieChannel`: as random zombies as generated, they are sent to
    one of 4 zombie channels, each representing the general direction
    from which a zombie spawned (north, south, east or west).
  * `HeroChannel`: every time a new zombie arrives, it is directed to a
    hero channel, according to their type (the weapon they carry). There
    are a total of 4 hero channels (see table above).
  * Technically, there's a third type of channel that is used for
    controlling the termination of the simulation.

### General requirements and notes

  * The number of zombies spawning from each direction is currently set
    in code.
  * Zombie names are of the form `D-NNNNN`, where `D` represents the
    source direction and `NNNNN` is a sequential number. E.g. the third
    zombie coming from the west will be named `W-00003`.
  * There should be a random delay between spawns.
  * Heroes can only fight one zombie at a time.
  * The hero channels are buffered and currently can hold up to 5
    zombies. If the zombie spawn rate is much higher than the time
    needed for a hero to defeat a zombie, blocking will occur and the
    simulation will last longer. Anyway, the fact is that zombies will
    be waiting in line before they are confronted by their respective
    hero. Yes, not very exciting, but that's how it is for now.
  * The simulation log includes 3 types of events: `SPAWN`, `FIGHT` and
    `KILL`.
  * All simulations should start with one or more `SPAWN` events.
  * Ideally, simulations should end with a `KILL` event. However, it
    might be possible that the *timeout since last spawn* is reached
    before all existing zombies have been defeated.

### Improvement ideas

  * External configuration of the simulation parameters.
    - Zombie spawn rate
    - Number of zombies spawning from each direction
    - Timeout from last spawn before the simulation is terminated.
  * Statistics summary at the end of simulations.

----
Andre Zunino &lt;neyzunino@gmail.com&gt;  
October 2020

