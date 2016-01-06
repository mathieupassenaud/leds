package backend

import "C"
import "github.com/mathieupassenaud/leds/ws2811"

func ApplyAll(statuses []statuses) {

	for i := 0; i < len(statuses); i++ {
		apply(statuses[i])
	}

	ws2811.Render()
}

func Apply(s status) {
	s.ledIndex
	s.currentColor
	s.mode
	s.frequency

	// TODAY We only apply colors
	value := colorToInt(s.currentColor.red, s.currentColor.green, s.currentColor.blue)

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
	return (red << 16) | (green << 8) | blue
}

func Init() {
	ws2811.Init(18, 1500, 255)
}
