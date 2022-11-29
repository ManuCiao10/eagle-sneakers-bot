package exe

import (
	"errors"
	"os/exec"

	"github.com/eagle/handler/version"
)

func Execute() {
	path, err := exec.LookPath("bin/EagleBot_" + version.Version + ".exe")
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(path)
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	if err := cmd.Run(); err != nil {
		panic(err)
	}

}
