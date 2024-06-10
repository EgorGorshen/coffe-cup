package particles

import (
	"math"
	"math/rand"
)

type Coffee struct {
	ParticlesSys
}

func normalize(cord int) float64 {
	mx := math.Floor(float64(cord) / 2)
	return math.Max(-mx, math.Min(rand.NormFloat64(), mx)) + mx
}

func reset(particle *Particle, params *ParticlesParams) {
    particle.lifetime = int64(math.Floor(float64(params.maxLife) * rand.Float64()))
    particle.speed = math.Floor(params.maxSpeed * rand.Float64())

    maxX := math.Floor(float64(params.X) / 2)
    particle.X = math.Max(-maxX, math.Min(rand.NormFloat64() * params.XScale, maxX)) + maxX
    particle.Y = 0
}

func ascii(row, col int, counts [][]int) string {
	count := counts[row][col]
	if count < 3 {
		return " "
	}
	if count < 6 {
		return "."
	}
	if count < 9 {
		return ":"
	}
	if count < 12 {
		return "}"
	}

	return "{"
}

func nextPostion(particle *Particle, deltaMS int64) {
	particle.lifetime -= deltaMS

	if particle.lifetime <= 0 {
		return
	}

	particle.Y += particle.speed * float64(deltaMS) / 1000.0
}

func NewCoffee(width, hight int, scale float64) Coffee {

    ascii := func(row, col int, counts [][]int) string {
           count := counts[row][col]
           if count == 0 {
               return " "
           }
           if count < 4 {
               return "░"
           }
           if count  < 6 {
               return "▒"
           }
           if count < 9 {
               return "▓"
           }
           return "█"
   }

	return Coffee{
		ParticlesSys: NewParticlesSys(ParticlesParams{
			maxLife:        7000,
			maxSpeed:       1.75 * 3,
			ParticlesCount: 1000,

			X: width,
			Y: hight,
            XScale: scale,

			reset:       reset,
			ascii:       ascii,
			nextPostion: nextPostion,
		}),
	}
}
