// Copyright (c) 2022 Daniel Krajka, All rights reserved.
// burl source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

// Package burl provides an easy, fluent way of building URLs via chainable methods
package burl

import (
	"net/url"

	"github.com/mr-cryka/burl/internal/str"
)

type Burl struct {
	Url         *url.URL
	pathParts   pathParts
	queryParams queryParams
}

// FromString creates a new Burl based on the url string
func FromString(urlString string) *Burl {
	if str.IsEmptyOrWhitespace(urlString) {
		return &Burl{
			Url: &url.URL{},
		}
	}

	res, err := url.Parse(urlString)
	if err != nil {
		return &Burl{
			Url: &url.URL{},
		}
	}

	return FromUrl(res)
}

// FromString creates a new Burl based on the (net/url) URL pointer
func FromUrl(url *url.URL) *Burl {
	return &Burl{
		Url:         url,
		queryParams: parseQueryParams(url.RawQuery),
		pathParts:   parsePathParts(url.Path),
	}
}

func (u *Burl) String() string {
	return u.Url.String()
}
