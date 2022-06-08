package internal

type internal struct {
	directoryRoot *string
	directory     string
}

func New() *internal {
	return &internal{}
}
func NewDirectoryRoot(directoryRoot string) *internal {
	return &internal{
		directoryRoot: &directoryRoot,
	}
}
