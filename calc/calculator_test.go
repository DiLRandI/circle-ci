package calc_test

import (
	"testing"

	"github.com/DiLRandI/circle-ci/calc"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	sut := calc.New()
	act := sut.Add(1, 1)

	assert.Equal(t, 2, act)
}
