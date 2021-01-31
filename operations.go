package gorpicam

import "fmt"

func pushBool(arr []string, name string, val bool) []string {

	if val {
		return append(arr, fmt.Sprintf("--%s", name))
	}
	return arr
}

func pushString(arr []string, name, val string) []string {
	return append(arr, fmt.Sprintf("--%s", name), val)
}

func pushInt8(arr []string, name string, val int8) []string {
	return pushString(arr, name, fmt.Sprintf("%d", val))
}

func pushUInt8(arr []string, name string, val uint8) []string {
	return pushString(arr, name, fmt.Sprintf("%d", val))
}

func pushInt16(arr []string, name string, val uint16) []string {
	return pushString(arr, name, fmt.Sprintf("%d", val))
}

func pushUInt16(arr []string, name string, val uint16) []string {
	return pushString(arr, name, fmt.Sprintf("%d", val))
}

func pushUInt(arr []string, name string, val uint) []string {
	return pushString(arr, name, fmt.Sprintf("%d", val))
}

func pushInt(arr []string, name string, val int) []string {
	return pushString(arr, name, fmt.Sprintf("%d", val))
}

func clampInt8(val, min, max int8) int8 {

	if min > max {
		panic(fmt.Sprintf("Min (%d) must be less than or equal to max (%d)", min, max))
	}

	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}

func clampUInt8(val, min, max uint8) uint8 {

	if min > max {
		panic(fmt.Sprintf("Min (%d) must be less than or equal to max (%d)", min, max))
	}

	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}

func clampInt16(val, min, max int16) int16 {

	if min > max {
		panic(fmt.Sprintf("Min (%d) must be less than or equal to max (%d)", min, max))
	}

	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}

func clampUInt16(val, min, max uint16) uint16 {

	if min > max {
		panic(fmt.Sprintf("Min (%d) must be less than or equal to max (%d)", min, max))
	}

	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}

func clampUInt(val, min, max uint) uint {

	if min > max {
		panic(fmt.Sprintf("Min (%d) must be less than or equal to max (%d)", min, max))
	}

	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}
