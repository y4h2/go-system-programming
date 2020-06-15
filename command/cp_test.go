package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCp(t *testing.T) {
	err := cp("./cp_test.go", "./temp.go")
	assert.NoError(t, err)
}
