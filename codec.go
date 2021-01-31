package gorpicam

// Codec specifies video codec used
type Codec string

const (
	// CodecDefault default encoding
	CodecDefault Codec = CodecH264

	// CodecH264 H264 encoding
	CodecH264 Codec = "H264"

	// MotionJPEG Motion JPEG encoding
	MotionJPEG Codec = "MJPEG"
)
