package ccutil

import (
	"errors"
	"os"
	"os/exec"
)

func CC() (string, error) {
	cands := []string{
		os.Getenv("CLANG"),
		os.Getenv("CC"),
		// clang seems better for a large C source
		"clang",
		"cc",
		"gcc",
	}
	for _, f := range cands {
		if f == "" {
			continue
		}
		exe, err := exec.LookPath(f)
		if err == nil {
			return exe, nil
		}
	}

	return "", errors.New("no C compiler was found")
}
