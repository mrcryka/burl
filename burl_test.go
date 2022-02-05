package burl

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FromString(t *testing.T) {
	// Arrange
	input := "https://www.test.com/foo?x=1&y=2"
	parsed, _ := url.Parse(input)

	// Act
	burl := FromString(input)

	// Assert
	assert.Equal(t, parsed.String(), burl.String())
}

func Test_FromString_Invalid_String_Returns_Empty(t *testing.T) {
	scenarios := []struct {
		name string
		input string
	}{
		{"Empty", ""},
		{"Whitespace", "   "},
		{"Invalid_Url", ";:;:;"},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			// Act
			burl := FromString(scenario.input)
		
			// Assert
			assert.Equal(t, "", burl.String())
		})
	}

}