package backend

import "C"
import (
	"github.com/mathieupassenaud/leds/ws2811"
)

func ApplyAll(statuses []Status) {

	for i := 0; i < len(statuses); i++ {
		Apply(statuses[i])
	}

	ws2811.Render()
}

func Apply(s Status) {
	//s.ledIndex
	//s.currentColor
	//s.mode
	//s.frequency

	// TODAY We only apply colors
	value := colorToInt(s.currentColor.Red, s.currentColor.Green, s.currentColor.Blue)

	ws2811.SetLed(s.ledIndex, value)

	// TODO create cases for :
	// blink(led, frequency)
	// rainbow(from, to)
	// change(from, to, colors[])

}

func ForceRender() {
	ws2811.Render()
}

func colorToInt(red int, green int, blue int) uint32 {
	var value uint32
	value = uint32((red << 16) | (green << 8) | blue)
	return value
}

func Init() {
	ws2811.Init(18, 1500, 255)
}
