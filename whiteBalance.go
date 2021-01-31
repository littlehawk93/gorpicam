package gorpicam

// WhiteBalance controls modes for temperature ranges
type WhiteBalance string

const (
	// WhiteBalanceDefault the default white balance setting
	WhiteBalanceDefault WhiteBalance = WhiteBalanceAuto

	// WhiteBalanceOff turn off white balance calculation
	WhiteBalanceOff WhiteBalance = "off"

	// WhiteBalanceAuto automatic mode
	WhiteBalanceAuto WhiteBalance = "auto"

	// WhiteBalanceSun sunny mode (between 5000K and 6500K)
	WhiteBalanceSun WhiteBalance = "sun"

	// WhiteBalanceCloud cloudy mode (between 6500K and 12000K)
	WhiteBalanceCloud WhiteBalance = "cloud"

	// WhiteBalanceShade shade mode
	WhiteBalanceShade WhiteBalance = "shade"

	// WhiteBalanceTungsten tungsten lighting mode (between 2500K and 3500K)
	WhiteBalanceTungsten WhiteBalance = "tungsten"

	// WhiteBalanceFlourescent fluorescent lighting mode (between 2500K and 4500K)
	WhiteBalanceFlourescent WhiteBalance = "flourescent"

	// WhiteBalanceIncandescent incandescent lighting mode
	WhiteBalanceIncandescent WhiteBalance = "incandescent"

	// WhiteBalanceFlash flash mode
	WhiteBalanceFlash WhiteBalance = "flash"

	// WhiteBalanceHorizon horizon mode
	WhiteBalanceHorizon WhiteBalance = "horizon"

	// WhiteBalanceGreyworld Use this on the NoIR camera to fix incorrect AWB results due to the lack of the IR filter.
	WhiteBalanceGreyworld WhiteBalance = "greyworld"
)
