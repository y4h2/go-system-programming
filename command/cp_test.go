package command

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCp(t *testing.T) {
	assert := assert.New(t)
	testingDir := ""
	testingDir, err := ioutil.TempDir(testingDir, "cp")
	assert.NoError(err)
	defer os.RemoveAll(testingDir)

	source := "./cp_test.go"
	target := testingDir + "/temp.go"
	err = cp(source, target)
	assert.NoError(err)
	assert.True(compareFiles(t, source, target))
}

func compareFiles(t *testing.T, f1, f2 string) bool {
	t.Helper()
	b1, err := ioutil.ReadFile(f1)
	if err != nil {
		t.Fatal(err)
	}

	b2, err := ioutil.ReadFile(f2)
	if err != nil {
		t.Fatal(err)
	}

	return bytes.Equal(b1, b2)
}
