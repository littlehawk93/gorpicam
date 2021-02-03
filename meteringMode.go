package gorpicam

// MeteringMode defines the kind of metering mode the camera performs for setting exposure
type MeteringMode string

const (
	// MeteringDefault default metering mode
	MeteringDefault MeteringMode = MeteringAverage

	// MeteringAverage average values across image
	MeteringAverage MeteringMode = "average"

	// MeteringSpot focus on central spot of image
	MeteringSpot MeteringMode = "spot"

	// MeteringBacklit meters for backlit subjects
	MeteringBacklit MeteringMode = "backlit"

	// MeteringMatrix uses a matrix of metering spots
	MeteringMatrix MeteringMode = "matrix"
)
