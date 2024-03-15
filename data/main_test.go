package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getLocale(t *testing.T) {
	assert.Equal(t, getLocale("zh-CN"), "zh")
	assert.Equal(t, getLocale("zh-TW"), "zh")
	assert.Equal(t, getLocale("en-US"), "en")
	assert.Equal(t, getLocale("other"), "en")
}
