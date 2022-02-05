package burl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetQueryParam(t *testing.T) {
	// Arrange
	expected := "http://www.test.com/test?x=1&y=2&y=3&y=4&z=5&abc&def&foo=&=bar"

	// Act
	burl := FromString("http://www.test.com/test").
		SetQueryParam("x", "1").
		SetQueryParamSlice("y", "2", "3", "4").
		SetQueryParam("z", "5").
		SetQueryParamKey("abc").
		SetQueryParamKey("def").
		SetQueryParam("foo", "").
		SetQueryParam("", "bar")

	// Assert
	assert.Equal(t, expected, burl.String())
}

func Test_SetQueryParamSlice_Overwrites_Existing_Param_Slice(t *testing.T) {
	// Arrange
	baseUrl := "http://www.test.com/test?x=1&x=2&x=3"
	expected := "http://www.test.com/test?x=100&x=200&x=300"

	// Act
	burl := FromString(baseUrl).
		SetQueryParamSlice("x", "100", "200", "300")

	// Assert
	assert.Equal(t, expected, burl.String())
}

func Test_SetQueryParams(t *testing.T) {
	scenarios := []struct {
		name     string
		params   []QueryParam
		expected string
	}{
		{"", []QueryParam{{"foo", "123"}, {"bar", "456"}}, "http://www.test.com/test?foo=123&bar=456"},
		{"Overwrites", []QueryParam{{"foo", "1"}, {"bar", "2"}, {"foo", "100"}}, "http://www.test.com/test?foo=100&bar=2"},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			// Arrange
			baseUrl := "http://www.test.com/test"

			// Act
			burl := FromString(baseUrl).
				SetQueryParams(scenario.params)

			// Assert
			assert.Equal(t, scenario.expected, burl.String())
		})
	}
}

func Test_SetQueryParamKey_Ignores_Empty_Query_Param_Key(t *testing.T) {
	// Arrange
	expected := "http://www.test.com/test"

	// Act
	burl := FromString(expected).SetQueryParamKey("")

	// Assert
	assert.Equal(t, expected, burl.String())
}

func Test_SetQueryParam_Overwrites_Existing_Query_Param(t *testing.T) {
	// Arrange
	expected := "http://www.test.com/test?hello=there"

	// Act
	burl := FromString("http://www.test.com/test?hello=world").
		SetQueryParam("hello", "there")

	// Assert
	assert.Equal(t, expected, burl.String())
}

func Test_SetQuery(t *testing.T) {
	scenarios := []struct {
		name     string
		query    string
		expected string
	}{
		{"", "x=1&y=2&z=3", "http://www.test.com/test?x=1&y=2&z=3"},
		{"Already_Encoded", "foo=space%20here", "http://www.test.com/test?foo=space%20here"},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			// Arrange
			baseUrl := "http://www.test.com/test"

			// Act
			burl := FromString(baseUrl).SetQuery(scenario.query)

			// Assert
			assert.Equal(t, scenario.expected, burl.String())
		})
	}

}

func Test_RemoveQueryParam(t *testing.T) {
	scenarios := []struct {
		name     string
		query    string
		remove   string
		expected string
	}{
		{"", "x=1&y=2&z=3", "y", "http://www.test.com/test?x=1&z=3"},
		{"Non_Existing_Query_Param", "x=1&y=2&z=3", "foo", "http://www.test.com/test?x=1&y=2&z=3"},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			// Arrange
			baseUrl := "http://www.test.com/test"

			// Act
			burl := FromString(baseUrl).
				SetQuery(scenario.query).
				RemoveQueryParam(scenario.remove)

			// Assert
			assert.Equal(t, scenario.expected, burl.String())
		})
	}
}
