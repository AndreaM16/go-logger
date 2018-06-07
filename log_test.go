package log

import (
	"testing"
	"bytes"
	"os"
	"github.com/stretchr/testify/assert"
)

func init() {
	Init(os.Stdout, os.Stdout, os.Stderr)
}

func TestWarning(t *testing.T) {

	buf := new(bytes.Buffer)
	Warning.SetOutput(buf)

	Warning.Print("")

	str := buf.String()

	assert.Contains(t, str, "WARNING")

	buf.Reset()

}

func TestInfo(t *testing.T) {

	buf := new(bytes.Buffer)
	Info.SetOutput(buf)

	Info.Print("")

	str := buf.String()

	assert.Contains(t, str, "INFO")

	buf.Reset()

}

func TestError(t *testing.T) {

	buf := new(bytes.Buffer)
	Error.SetOutput(buf)

	Error.Print("")

	str := buf.String()

	assert.Contains(t, str, "ERROR")

	buf.Reset()

}

