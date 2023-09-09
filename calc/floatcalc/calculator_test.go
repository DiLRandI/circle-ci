package floatcalc_test

import (
	"testing"

	"github.com/DiLRandI/circle-ci/calc/floatcalc"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	sut := floatcalc.New()

	act := sut.Add(1, 1)

	assert.Equal(t, 2.0, act)
}

func TestAddFloat(t *testing.T) {
	t.Parallel()

	sut := floatcalc.New()

	act := sut.Add(1.5, 1.5)

	assert.Equal(t, 3.0, act)
}
