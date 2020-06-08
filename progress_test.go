package logutil

import (
	_ "github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProgress(t *testing.T) {
	p := NewProgress(LoggerFunc(t.Logf), "TEST")
	p.SetCount(10)
	p.SetTotal(200)
	for i := 0; i < 200; i++ {
		p.Incr()
	}
	p.Add(2)
	p.SetCount(100)

	p = NewProgress(nil, "TEST")
	p.SetCount(10)
	p.SetTotal(200)
	for i := 0; i < 200; i++ {
		p.Incr()
	}
	p.Add(2)
	p.SetCount(100)
}
