package parse

import (
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// Module returns the module name in the current directory
func Module() (string, error) {
	out, err := exec.Command("go.exe", "list", "-m").Output()
	if err != nil {
		return "", errors.Wrap(err, "exec.command")
	}
	str := string(out)
	str = strings.TrimSpace(str)
	return str, nil
}
