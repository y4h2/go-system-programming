package fileio

import (
	"bufio"
	"os"

	"github.com/pkg/errors"
)

/*
相比io package, bufio更灵活, 允许line-by-line的获取

*/

func bufferRead(filename string) (result []string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return result, errors.Wrapf(err, "failed to open %s", filename)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		err = scanner.Err()
		if err != nil {
			return result, errors.Wrapf(err, "failed to read %s", filename)
		}

		result = append(result, line)
	}
	return result, nil
}
