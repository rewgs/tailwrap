package tailwrap

import (
	"os"
	"os/exec"
	"runtime"
)

type Tailscale struct {
	Bin     string
	Variant string // gui or cli; make this an enum?
}

// getBin looks for tailscale
func getBin() string {
	if runtime.GOOS == "darwin" {
		os.Open("/Applications/Tailscale.app/Contents/MacOS/Tailscale")
	}
}
