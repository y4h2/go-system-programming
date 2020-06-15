package command

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

const BUF_SIZE = 1024

func cp(src, dest string) error {
	iFile, err := os.Open(src)
	if err != nil {
		return errors.Wrap(err, "open input file")
	}

	oFile, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return errors.Wrap(err, "open output fiel")
	}

	b := make([]byte, BUF_SIZE)
	for {
		n, err := iFile.Read(b)
		if n == 0 && err == io.EOF { // reach the end of file
			break
		}
		if n < BUF_SIZE { // avoid NUL characters
			b = b[:n]
		}

		n, err = oFile.Write(b)
		if len(b) != n {
			errors.New("couldn't write whole buffer")
		}
		if err != nil {
			return err
		}
	}

	err = oFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
