package burl

import (
	"burl/internal/str"
	"net/url"
	"strings"
)

type pathParts struct {
	parts   []string
	indexes map[string]int
}

// SetPath sets an entire url path
func (u *Burl) SetPath(path string) *Burl {
	if str.IsEmptyOrWhitespace(path) {
		return u
	}

	u.pathParts = parsePathParts(path)
	u.rebuildPath()

	return u
}

// AddPathPart adds a single path part
func (u *Burl) AddPathPart(path string) *Burl {
	if str.IsEmptyOrWhitespace(path) {
		return u
	}

	u.pathParts.parts = append(u.pathParts.parts, path)
	u.rebuildPath()

	return u
}

// AddPathPart adds multiple path parts
func (u *Burl) AddPathParts(paths ...string) *Burl {
	for _, path := range paths {
		if str.IsEmptyOrWhitespace(path) {
			continue
		}

		u.pathParts.parts = append(u.pathParts.parts, path)
	}

	u.rebuildPath()

	return u
}

// AddPathPart removes a single path part
func (u *Burl) RemovePathPart(path string) *Burl {
	index, exists := u.pathParts.indexes[path]
	if !exists {
		return u
	}

	u.pathParts.parts = append(u.pathParts.parts[:index], u.pathParts.parts[index+1:]...)
	u.rebuildPath()

	return u
}

func (u *Burl) rebuildPath() {
	joined := strings.Join(u.pathParts.parts, "/")
	parsedJoined, _ := url.Parse(joined)
	u.Url = u.Url.ResolveReference(parsedJoined)
}

func parsePathParts(pathString string) pathParts {
	parts := strings.Split(pathString, "/")

	result := pathParts{
		indexes: make(map[string]int),
	}

	for i, part := range parts {
		result.parts = append(result.parts, part)
		result.indexes[part] = i
	}

	return result
}
