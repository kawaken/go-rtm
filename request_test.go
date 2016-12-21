package rtm

import (
	"testing"
)

func TestSign(t *testing.T) {

	params := newParams()

	s := sign(params)

	t.Log(s)
}
