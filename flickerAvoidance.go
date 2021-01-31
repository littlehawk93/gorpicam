package gorpicam

// FlickerAvoidance mode to compensate for flickering lights
type FlickerAvoidance string

const (
	// FlickerOff turn off flicker avoidance
	FlickerOff FlickerAvoidance = "off"

	// FlickerAuto automatically detect frequency
	FlickerAuto FlickerAvoidance = "auto"

	// Flicker50 avoidance for 50 Hz lights
	Flicker50 FlickerAvoidance = "50hz"

	// Flicker60 avoidance for 60 Hz lights
	Flicker60 FlickerAvoidance = "60hz"
)
