package burl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddPathPart(t *testing.T) {
	scenarios := []struct {
		name     string
		path1    string
		path2    string
		expected string
	}{
		{"Full_Url", "http://www.test.com", "test", "http://www.test.com/test"},
		{"Paths", "foo", "bar", "/foo/bar"},
		{"Relative", "/foo/bar", "xyz", "/foo/bar/xyz"},
		{"Empty_Path", "http://www.test.com", "", "http://www.test.com"},
		{"Already_Encoded", "http://www.test.com", "path%20with%20spaces", "http://www.test.com/path%20with%20spaces"},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			// Arrange

			// Act
			burl := FromString(scenario.path1).
				AddPathPart(scenario.path2)

			// Assert
			assert.Equal(t, scenario.expected, burl.String())
		})
	}
}

func Test_AddPathParts(t *testing.T) {
	// Arrange
	baseUrl := "http://www.test.com"
	paths := []string{"foo with spaces", "", "bar"}

	// Act
	burl := FromString(baseUrl).
		AddPathParts(paths...)

	// Assert
	assert.Equal(t, "http://www.test.com/foo%20with%20spaces/bar", burl.String())
}

func Test_RemovePathPart_Removes_Path_Part(t *testing.T) {
	// Arrange
	scenarios := []struct {
		input        string
		pathToRemove string
		expected     string
	}{
		{"/path1/path2", "path1", "/path2"},
		{"/foo/bar", "bar", "/foo"},
		{"/foo/bar", "xyz", "/foo/bar"},
	}

	for _, scenario := range scenarios {
		// Act
		burl := FromString(scenario.input).
			RemovePathPart(scenario.pathToRemove)

		// Assert
		assert.Equal(t, scenario.expected, burl.String())
	}

}

func Test_SetPath(t *testing.T) {
	scenarios := []struct {
		name     string
		baseUrl  string
		path     string
		expected string
	}{
		{"Sets_Path", "http://test.com/", "/foo/bar", "http://test.com/foo/bar"},
		{"With_Space", "http://test.com", "/this/is/some path", "http://test.com/this/is/some%20path"},
		{"Already_Encoded", "http://test.com", "/this/is/some%20path", "http://test.com/this/is/some%20path"},
		{"Empty_Path", "http://test.com", "", "http://test.com"},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			// Act
			burl := FromString(scenario.baseUrl).
				SetPath(scenario.path)

			// Arrange
			assert.Equal(t, scenario.expected, burl.String())
		})
	}

}
