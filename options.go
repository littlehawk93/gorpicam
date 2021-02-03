package gorpicam

import "fmt"

// Options generic options for any raspicam command
type Options struct {
	Width          uint16           `mapstructure:"width" json:"width"`
	Height         uint16           `mapstructure:"height" json:"height"`
	FlipHorizontal bool             `mapstructure:"hor_flip" json:"hor_flip"`
	FlipVertical   bool             `mapstructure:"ver_flip" json:"ver_flip"`
	Preview        []uint16         `mapstructure:"preview" json:"preview"`
	NoPreview      bool             `mapstructure:"nppreview" json:"nopreview"`
	FullScreen     bool             `mapstructure:"preview_fullscreen" json:"preview_fullscreen"`
	PreviewOpacity uint8            `mapstructure:"preview_opacity" json:"preview_opacity"`
	Sharpness      int8             `mapstructure:"sharpness" json:"mapstructure"`
	Contrast       int8             `mapstructure:"contrast" json:"contrast"`
	Brightness     uint8            `mapstructure:"brightness" json:"brightness"`
	Saturation     int8             `mapstructure:"saturation" json:"saturation"`
	ISO            uint16           `mapstructure:"iso" json:"iso"`
	EV             int8             `mapstructure:"ev" json:"ev"`
	Exposure       ExposureMode     `mapstructure:"exposure" json:"exposure"`
	Flicker        FlickerAvoidance `mapstructure:"flicker" json:"flicker"`
	WhiteBalance   WhiteBalance     `mapstucture:"white_balance" json:"white_balance"`
	Metering       MeteringMode     `mapstructure:"metering" json:"metering"`
	Annotate       string           `mapstructure:"annotate" json:"annotate"`
}

// CmdArgs get the command line arguments for executign a raspicam command
func (me Options) CmdArgs() []string {

	args := make([]string, 0)

	args = pushUInt16(args, "width", clampUInt16(me.Width, 1, 4096))
	args = pushUInt16(args, "height", clampUInt16(me.Height, 1, 2160))

	args = pushBool(args, "hflip", me.FlipHorizontal)
	args = pushBool(args, "vflip", me.FlipVertical)

	if me.NoPreview {
		args = pushBool(args, "preview", me.NoPreview)
	} else if me.FullScreen {
		args = pushBool(args, "fullscreen", me.FullScreen)

		if me.PreviewOpacity < 255 {
			args = pushUInt8(args, "opacity", me.PreviewOpacity)
		}
	} else if me.Preview != nil && len(me.Preview) == 4 {

		args = pushString(args, "preview", fmt.Sprintf("%d,%d,%d,%d", me.Preview[0], me.Preview[1], me.Preview[2], me.Preview[3]))

		if me.PreviewOpacity < 255 {
			args = pushUInt8(args, "opacity", me.PreviewOpacity)
		}
	}

	if me.Sharpness != 0 {
		args = pushInt8(args, "sharpness", clampInt8(me.Sharpness, -100, 100))
	}

	if me.Contrast != 0 {
		args = pushInt8(args, "contrast", clampInt8(me.Contrast, -100, 100))
	}

	if me.Brightness != 0 {
		args = pushUInt8(args, "brightness", clampUInt8(me.Brightness, 0, 100))
	}

	if me.Saturation != 0 {
		args = pushInt8(args, "saturation", clampInt8(me.Saturation, -100, 100))
	}

	args = pushUInt16(args, "iso", clampUInt16(me.ISO, 100, 800))

	if me.EV != 0 {
		args = pushInt8(args, "ev", clampInt8(me.EV, -10, 10))
	}

	if me.Exposure != ExposureDefault {
		args = pushString(args, "exposure", string(me.Exposure))
	}

	if me.Flicker != FlickerOff {
		args = pushString(args, "flicker", string(me.Flicker))
	}

	if me.WhiteBalance != WhiteBalanceDefault {
		args = pushString(args, "awb", string(me.WhiteBalance))
	}

	if me.Metering != MeteringDefault {
		args = pushString(args, "metering", string(me.Metering))
	}

	if me.Annotate != "" {
		args = pushString(args, "annotate", me.Annotate)
	}

	return args
}

func defaultOptions() Options {

	return Options{
		Width:          1920,
		Height:         1080,
		FlipHorizontal: false,
		FlipVertical:   false,
		NoPreview:      true,
		Preview:        nil,
		FullScreen:     false,
		PreviewOpacity: 100,
		Sharpness:      0,
		Contrast:       0,
		Brightness:     0,
		Saturation:     0,
		ISO:            200,
		EV:             0,
		Exposure:       ExposureDefault,
		Flicker:        FlickerOff,
		WhiteBalance:   WhiteBalanceDefault,
		Metering:       MeteringDefault,
		Annotate:       "",
	}
}
