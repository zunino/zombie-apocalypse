# Zombie Apocalypse

This program was inspired by a [zombie apocalypse problem statement created by Johni](https://gist.github.com/johnidm/3e98e4ff7565ce985a41).
His original intent was to use the scenario as the basis for playing
with golang's concurrent concepts and primitives.

In Johni's original conception, zombies use *1 of 4 different types
of weapons* and there are *4 heroes* who are supposed to defeat the
invading zombies. Each hero specializes in one kind of zombie weapon,
meaning they will only fight those carrying the appropriate weapons. 

| Zombie weapon    | Hero       |
|------------------|------------|
| Baseball bat     | Tex Willer |
| Hatchet          | Kit Carson |
| Scythe           | Jack Tiger |
| Machete          | Kit Willer |

## Adaptations

In the original design, channels were used to represent each type of
zombie. In the present design, channels represent the direction zombies
are spawning from: *north*, *south*, *east* and *west*.

----
Andre Zunino &lt;neyzunino@gmail.com&gt;  
October 2020

