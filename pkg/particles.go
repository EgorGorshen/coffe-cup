package particles

import (
	"math"
	"slices"
	"strings"
	"time"
)

type Particle struct {
	lifetime int64
	speed    float64
	X        float64
	Y        float64
}

type ParticlesParams struct {
	maxLife        float64
	maxSpeed       float64
	ParticlesCount int
	X              int
    XScale float64
	Y              int
	nextPostion    NextPostionFunk
	ascii          ASCII
	reset          Reset
}

type NextPostionFunk = func(particle *Particle, deltaMS int64)
type ASCII func(x, y int, count [][]int) string
type Reset func(particle *Particle, params *ParticlesParams)

type ParticlesSys struct {
	ParticlesParams
	LastTime  int64
	particles []*Particle
}

func NewParticlesSys(params ParticlesParams) ParticlesSys {
	particles := make([]*Particle, 0)
	for i := 0; i < params.ParticlesCount; i++ {
		particles = append(particles, &Particle{})
	}
	return ParticlesSys{
		ParticlesParams: params,
		LastTime:        time.Now().UnixMilli(),
		particles:       particles,
	}
}

func (self *ParticlesSys) Update() {
	now := time.Now().UnixMilli()
	delta := now - self.LastTime
	self.LastTime = now
	for _, p := range self.particles {
		self.nextPostion(p, delta)
		if p.X >= float64(self.X) || p.Y >= float64(self.Y) || p.lifetime <= 0 {
			self.reset(p, &self.ParticlesParams)
		}
	}
}

func (self *ParticlesSys) Display() string {
	counts := make([][]int, 0)
	for row := 0; row < self.Y; row++ {
		count := make([]int, 0)
		for col := 0; col < self.X; col++ {
			count = append(count, 0)
		}
		counts = append(counts, count)
	}

	for _, p := range self.particles {
		row := int(math.Floor(p.Y))
		col := int(math.Floor(p.X))
		counts[row][col]++
	}
	out := make([][]string, 0)
	for r, row := range counts {
		outRow := make([]string, 0)
		for c := range row {
			outRow = append(outRow, self.ascii(r, c, counts))
		}
		out = append(out, outRow)
	}

	slices.Reverse(out)

	outStr := make([]string, 0)

	for _, row := range out {
		outStr = append(outStr, strings.Join(row, ""))
	}

	return strings.Join(outStr, "\n")
}

func (self *ParticlesSys) Start() {
	for _, p := range self.particles {
		self.reset(p, &self.ParticlesParams)
	}
}
