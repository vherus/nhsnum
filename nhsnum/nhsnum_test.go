package nhsnum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert.True(t, IsValid("5990128088"), "5990128088 is valid")
	assert.True(t, IsValid("1275988113"), "1275988113 is valid")
	assert.True(t, IsValid("4536026665"), "4536026665 is valid")
}

func TestIsNotValid(t *testing.T) {
	assert.False(t, IsValid("5990128087"), "5990128087 is not valid")
	assert.False(t, IsValid("4536016660"), "4536016660 is not valid")
}
