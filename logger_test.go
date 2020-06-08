package logutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoggerFunc_Printf(t *testing.T) {
	invoked := false
	var f LoggerFunc = func(layout string, items ...interface{}) {
		invoked = true
	}
	f.Printf("")
	assert.True(t, invoked)
}
