package fileio

import (
	"bytes"
	"encoding/binary"

	"github.com/pkg/errors"
)

func BigEndian(number int64) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, number)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate big endian")
	}

	return buf, nil
}

func LittleEndian(number int64) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, number)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate little endian")
	}

	return buf, nil
}
