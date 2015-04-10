package colors

// The colors package is directly based off of the
// colorsys python module.
// https://hg.python.org/cpython/file/2.7/Lib/colorsys.py

// TODO(aoeu): Refactor mulitplication and division by 360.

import (
	"fmt"
	"math"
)

const (
	oneThird  = 0.3333333333333333
	oneSixth  = 0.1666666666666666
	twoThirds = 0.6666666666666666
)

type HSL struct {
	Hue        float64
	Saturation float64
	Lightness  float64
}

func NewHSLSet(size int, saturation, lightness float64) []HSL {
	h := make([]HSL, size)
	for i := 0; i < size; i++ {
		hue := (360.0 / float64(size)) * float64(i)
		h[i] = HSL{hue, saturation, lightness}
	}
	return h
}

func (h HSL) ToRGB() RGB {
	if h.Saturation == 0 {
		return RGB{h.Lightness, h.Lightness, h.Lightness}
	}
	var m2 float64
	if h.Lightness <= 0.5 {
		m2 = h.Lightness * (1.0 + h.Saturation)
	} else {
		m2 = h.Lightness + h.Saturation - (h.Lightness * h.Saturation)
	}
	hue := h.Hue / 360.0
	m1 := 2.0*h.Lightness - m2
	return RGB{
		Red:   v(m1, m2, hue+oneThird),
		Green: v(m1, m2, hue),
		Blue:  v(m1, m2, hue-oneThird),
	}
}

func (h HSL) String() string {
	// TODO(aoeu): This should not be necessary.
	hue := int(h.Hue)
	if hue < 0 {
		hue = 0
	}
	return fmt.Sprintf("hsl(%v, %.0f%%, %.0f%%)",
		int(hue), (100.0 * h.Saturation), (100.0 * h.Lightness))
}

func v(m1, m2, hue float64) float64 {
	hue = math.Mod(hue, 1.0)
	switch {
	case hue < oneSixth:
		return m1 + (m2-m1)*hue*6.0
	case hue < 0.5:
		return m2
	case hue < twoThirds:
		return m1 + (m2-m1)*(twoThirds-hue)*6.0
	}
	return m1
}

type RGB struct {
	Red   float64
	Green float64
	Blue  float64
}

func (rgb RGB) String() string {
	var r, g, b uint8 = uint8(255 * rgb.Red), uint8(255 * rgb.Green), uint8(255 * rgb.Blue)
	return fmt.Sprintf("#%.2x%.2x%.2x", r, g, b)
}
