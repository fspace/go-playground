package ElegantConstants

const (
	one int = iota + 1
	two
)

const (
	KB ByteSize = 1 << (10 * (iota + 1))
	MB
	GB
	TB
	PB
)
