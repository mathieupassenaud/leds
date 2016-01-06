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
	Red   int
	Green int
	Blue  int
}

type Status struct {
	ledIndex     int
	currentColor Color
	mode         Mode
	frequency   int
}

var leds [1500]Status

func ChangeStatus(ledIndex int, setColor Color, setMode Mode, setFrequency int) Status {
	var led = leds[ledIndex]

	led.ledIndex = ledIndex
	led.currentColor = setColor
	led.mode = setMode
	led.frequency = setFrequency

	leds[ledIndex] = led
	return led
}

func GetStatusAlone(ledIndex int) Status {
	return leds[ledIndex]
}

func GetStatus(from int, to int) []Status {
	numberOfLeds := to - from
	var statuses []Status
	statuses = make([]Status, numberOfLeds)

	for i := 0; i < numberOfLeds; i = i + 1 {
		statuses[i] = leds[i+from]
	}

	return statuses
}
