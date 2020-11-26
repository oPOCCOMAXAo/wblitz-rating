package utils

import (
	"fmt"
	"math"
)

type ProgressBar struct {
	label    string
	total    float64
	progress float64
}

func NewProgressBar(label string, total float64) *ProgressBar {
	return &ProgressBar{
		label:    label,
		total:    total / 100,
		progress: 0,
	}
}

func (p *ProgressBar) log() {
	fmt.Printf("\r%s: %.2f%%", p.label, math.Min(p.progress/p.total, 100))
}

func (p *ProgressBar) Add(amount float64) {
	p.progress += amount
	p.log()
}

func (p *ProgressBar) Set(progress float64) {
	p.progress = progress
	p.log()
}
