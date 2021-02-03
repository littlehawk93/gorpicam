package gorpicam

const (
	// LengthUnlimited video recording doesn't stop until the process is killed
	LengthUnlimited uint = 0
)

// VidOptions contains various configuration options for a raspivid command
type VidOptions struct {
	Options            `mapstructure:"opt" json:"opt"`
	VideoStabilization bool  `mapstructure:"vstab" json:"vstab"`
	BitRate            uint  `mapstructure:"bps" json:"bps"`
	FrameRate          uint8 `mapstructure:"fps" json:"fps"`
	Codec              Codec `mapstructure:"codec" json:"codec"`
	Length             uint  `mapstructure:"length" json:"length"`
}

// CmdArgs get the command line arguments for executing a raspivid command
func (me VidOptions) CmdArgs() []string {

	args := me.Options.CmdArgs()

	args = pushUInt(args, "bitrate", clampUInt(me.BitRate, 0, 25000000))

	args = pushBool(args, "vstab", me.VideoStabilization)

	args = pushUInt8(args, "framerate", clampUInt8(me.FrameRate, 2, 30))

	args = pushString(args, "codec", string(me.Codec))

	args = pushUInt(args, "timeout", me.Length)
	return args
}

// DefaultVid returns the default options for a raspivid command
func DefaultVid() VidOptions {

	return VidOptions{
		Options:            defaultOptions(),
		BitRate:            17000000,
		VideoStabilization: false,
		FrameRate:          30,
		Codec:              CodecDefault,
		Length:             5000,
	}
}
