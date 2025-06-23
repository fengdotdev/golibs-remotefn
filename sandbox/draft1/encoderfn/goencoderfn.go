package encoderfn

type mode string

const (
	ModeJSON mode = "json"
	ModeGob  mode = "gob"
)

type GoEncoderFn struct {
	mode mode
}
