package backend

import ()

type Mode int

const (
	normal  Mode = 1
	rainbow Mode = 2
	change  Mode = 3
	blink   Mode = 4
)

type Color struct {
	red   int
	green int
	blue  int
}

type status struct {
	ledIndex     int
	currentColor color
	mode         Mode
	frenquency   int
}

var leds [1500]status

func ChangeStatus(ledIndex int, setColor color, setMode mode, setFrequency frequency) status {
	var led = leds[ledIndex]

	led.ledIndex = ledIndex
	led.currentColor = setColor
	led.mode = setMode
	led.frequency = setFrequency

	leds[ledIndex] = led
	return led
}

func getStatus(ledIndex int) status {
	return leds[ledIndex]
}

func getStatus(from int, to int) []status {
	numberOfLeds := to - from
	var statuses [numberOfLeds]status

	for i := 0; i < numberOfLeds; i = i + 1 {
		statuses[i] = leds[i+from]
	}

	return statuses
}
