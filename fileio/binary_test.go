package fileio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBigEndian(t *testing.T) {
	buf, err := BigEndian(int64(1))
	assert.NoError(t, err)
	t.Logf("Big Endian: %x", buf)
}

func TestLittleEndian(t *testing.T) {
	buf, err := LittleEndian(int64(1))
	assert.NoError(t, err)
	t.Logf("Little Endian: %x", buf)
}
