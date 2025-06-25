package playclient

import "github.com/fengdotdev/golibs-remotefn/cmd/playground/playserver"

func DataInOut(key string, data []byte) ([]byte, error) {
	// boundary client / server code
	exec := playserver.NewDataExec()
	outdata, err := exec.DataInOut(key, data)
	if err != nil {
		return nil, err
	}
	// boundary client / server code end
	return outdata, nil
}
