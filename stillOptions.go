package gorpicam

const (
	// TimelapseNone no timelapse for still images
	TimelapseNone int = -1
)

// StillOptions contains various configuration options for a raspistill command
type StillOptions struct {
	Options   `mapstructure:"opt" json:"opt"`
	Quality   uint8 `mapstructure:"quality" json:"quality"`
	Raw       bool  `mapstructure:"raw" json:"raw"`
	Timeout   uint  `mapstructure:"timeout" json:"timeout"`
	Timelapse int   `mapstructure:"timelapse" json:"timelapse"`
}

// CmdArgs get the command line arguments for executing a raspistill command
func (me StillOptions) CmdArgs() []string {

	args := me.Options.CmdArgs()

	args = pushUInt8(args, "quality", clampUInt8(me.Quality, 0, 100))

	args = pushBool(args, "raw", me.Raw)

	args = pushUInt(args, "timeout", me.Timeout)

	if me.Timelapse != TimelapseNone {
		args = pushInt(args, "timelapse", me.Timelapse)
	}
	return args
}

// DefaultStill returns the default options for a raspistill command
func DefaultStill() StillOptions {

	return StillOptions{
		Options:   defaultOptions(),
		Quality:   75,
		Raw:       false,
		Timeout:   5000,
		Timelapse: TimelapseNone,
	}
}
