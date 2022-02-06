# Burl - a fluent URL builder
Burl provides an easy, fluent way of building URLs via chainable methods.

## How to install
```bash
go get github.com/mr-cryka/burl
```

## Quick start
```go
package main

import (
	"burl"
	"fmt"
)

func main() {
	url := burl.
		FromString("http://example.com?remove_me=please").
		AddPathParts("foo", "bar").
		SetQueryParam("hello", "who is this?").
		RemoveQueryParam("remove_me")
	
	fmt.Println(url.String())
}
// Output: http://example.com/foo/bar?hello=who+is+this%3F
```

## Available operations
### From - create new burl
```go
	// FromString
	fromString := burl.FromString("http://example.com")

	// FromUrl (net/url)
	baseUrl, _ := url.Parse("http://example.com")
	fromUrl := burl.FromUrl(baseUrl)
```
### Path
```go
	result := burl.FromString("http://example.com").
		SetPath("/path1/path2").        // Set the entire path
		AddPathPart("path3").           // Add a single path part
		AddPathParts("path4", "path5"). // Add multiple path parts
		RemovePathPart("path2")         // Remove a single path part

	fmt.Println(result)

	// Output: http://example.com/path1/path3/path4/path5
```
### Query
All set methods overwrite any existing query param (take a closer look at the "x" query param)
```go
	result := burl.FromString("http://example.com").
		SetQuery("q=1").                                                        // Set the entire query
		SetQueryParam("x", "2").                                                // Set a single query param
		SetQueryParam("remove", "me").
		SetQueryParams([]burl.QueryParam{{"y", "3"}, {"z", "4"}, {"x", "0"}}).  // Set multiple query params, here we overwrite the "x" query param
		SetQueryParamKey("isDone").                                             // Set a query param without a value
		SetQueryParamSlice("colors", "red", "green", "blue").                   // Set an array query param
		RemoveQueryParam("remove")                                              // Remove a single query param

	fmt.Println(result)

	// Output: http://example.com?q=1&x=0&y=3&z=4&isDone&colors=red&colors=green&colors=blue
```


## License

MIT