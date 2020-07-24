package count_test

import (
	"github.com/bruteforce1414/test-travis/count"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateSuccess(t *testing.T) {
	assert.Equal(t, count.Count(), 2)
}

func TestCalculateFail(t *testing.T) {
	assert.Equal(t, count.Count(), 3)
}