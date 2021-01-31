package gorpicam

type MeteringMode string

const (
	// MeteringDefault default metering mode
	MeteringDefault MeteringMode = MeteringAverage

	MeteringAverage MeteringMode = "average"

	MeteringSpot MeteringMode = "spot"

	MeteringBacklit MeteringMode = "backlit"

	MeteringMatrix MeteringMode = "matrix"
)
