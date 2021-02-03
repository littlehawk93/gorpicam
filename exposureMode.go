package gorpicam

// ExposureMode raspicam exposure modes
type ExposureMode string

const (
	// ExposureDefault the default exposure mode setting
	ExposureDefault ExposureMode = ExposureAuto

	// ExposureAuto automatic exposure mode
	ExposureAuto ExposureMode = "auto"

	// ExposureNight night shooting
	ExposureNight ExposureMode = "night"

	// ExposureNightPreview night shooting
	ExposureNightPreview ExposureMode = "nightpreview"

	// ExposureBacklight for backlit subjects
	ExposureBacklight ExposureMode = "backlight"

	// ExposureSpotlight for spotlit subjects
	ExposureSpotlight ExposureMode = "spotlight"

	// ExposureSports setting for sports (fast shutter etc.)
	ExposureSports ExposureMode = "sports"

	// ExposureSnow optimised for snowy scenery
	ExposureSnow ExposureMode = "snow"

	// ExposureBeach optimised for beach
	ExposureBeach ExposureMode = "beach"

	// ExposureVeryLong for long exposures
	ExposureVeryLong ExposureMode = "verylong"

	// ExposureFixedFPS constrain fps to a fixed value
	ExposureFixedFPS ExposureMode = "fixedfps"

	// ExposureAntishake antishake mode
	ExposureAntishake ExposureMode = "antishake"

	// ExposureFireworks optimised for fireworks
	ExposureFireworks ExposureMode = "fireworks"
)
