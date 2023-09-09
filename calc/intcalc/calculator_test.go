package intcalc_test

import (
	"testing"

	"github.com/DiLRandI/circle-ci/calc/intcalc"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	sut := intcalc.New()

	act := sut.Add(1, 1)

	assert.Equal(t, 2, act)
}
