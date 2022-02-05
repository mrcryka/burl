package burl

import (
	"burl/internal/str"
	"net/url"
	"strings"
)

type QueryParam struct {
	Key   string
	Value string
}

type queryParam struct {
	key    string
	values []string
}

type queryParams struct {
	params  []queryParam
	indexes map[string]int
}

// SetQuery sets an entire query
func (u *Burl) SetQuery(query string) *Burl {
	u.Url.RawQuery = query
	u.queryParams = parseQueryParams(query)

	return u
}

// SetQueryParam sets a query param with a given key and value
func (u *Burl) SetQueryParam(key, value string) *Burl {
	queryParams := u.queryParams.params

	if index, exists := u.queryParams.indexes[key]; exists {
		u.queryParams.params[index].values = []string{value}
	} else {
		u.queryParams.params = append(queryParams, queryParam{key, []string{value}})
	}

	query := u.rebuildQuery()
	u.Url.RawQuery = query

	return u
}

// SetQueryParamKey sets an empty value query param
//
// e.g for a key "someFlag" the result is ...?someFlag&...
func (u *Burl) SetQueryParamKey(key string) *Burl {
	if str.IsEmptyOrWhitespace(key) {
		return u
	}

	queryParams := u.queryParams.params

	u.queryParams.params = append(queryParams, queryParam{key, nil})

	query := u.rebuildQuery()
	u.Url.RawQuery = query

	return u
}

// SetQueryParamSlice sets an array of values for a single query key
//
// e.g for a key "colors" and values "red", "green" 
// the result is ...?colors=red&colors=green&...
func (u *Burl) SetQueryParamSlice(key string, values ...string) *Burl {
	queryParams := u.queryParams.params

	if index, exists := u.queryParams.indexes[key]; exists {
		u.queryParams.params[index].values = values
	} else {
		u.queryParams.params = append(queryParams, queryParam{key, values})
	}

	query := u.rebuildQuery()
	u.Url.RawQuery = query

	return u
}

// SetQueryParams sets multipley query params with Key Value pairs
func (u *Burl) SetQueryParams(queryParams []QueryParam) *Burl {
	for i, qp := range queryParams {
		if index, exists := u.queryParams.indexes[qp.Key]; exists {
			u.queryParams.params[index].values = []string{qp.Value}
		} else {
			u.queryParams.params = append(u.queryParams.params, queryParam{qp.Key, []string{qp.Value}})
			u.queryParams.indexes[qp.Key] = i
		}
	}

	query := u.rebuildQuery()
	u.Url.RawQuery = query

	return u
}

// RemoveQueryParam removes a single query param
func (u *Burl) RemoveQueryParam(key string) *Burl {
	index, exists := u.queryParams.indexes[key]
	if !exists {
		return u
	}

	newSlice := append(u.queryParams.params[:index], u.queryParams.params[index+1:]...)
	u.queryParams.params = newSlice

	u.Url.RawQuery = u.rebuildQuery()

	return u
}

func (u *Burl) rebuildQuery() string {
	var buf strings.Builder

	u.queryParams.indexes = make(map[string]int)

	for i, queryParam := range u.queryParams.params {

		// AddQueryParamKey path
		if queryParam.values == nil {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}

			keyEscaped := url.QueryEscape(queryParam.key)
			buf.WriteString(keyEscaped)
			u.queryParams.indexes[queryParam.key] = i

			continue
		}

		for _, value := range queryParam.values {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}

			keyEscaped := url.QueryEscape(queryParam.key)
			valueEscaped := url.QueryEscape(value)

			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(valueEscaped)

			u.queryParams.indexes[queryParam.key] = i
		}
	}

	return buf.String()
}

func parseQueryParams(queryParamsString string) queryParams {
	params := strings.Split(queryParamsString, "&")

	queryParams := queryParams{
		indexes: map[string]int{},
	}
	for i, param := range params {
		if str.IsEmptyOrWhitespace(param) {
			continue
		}

		separatorIndex := strings.IndexRune(param, '=')
		if separatorIndex == -1 {
			continue
		}

		rawKey, rawValue := param[:separatorIndex], param[separatorIndex+1:]
		key, keyErr := url.QueryUnescape(rawKey)
		if keyErr != nil {
			// TODO not sure if we should continue or return err
			continue
		}

		value, valueErr := url.QueryUnescape(rawValue)
		if valueErr != nil {
			// TODO not sure if we should continue or return err
			continue
		}

		index, alreadyExists := queryParams.indexes[key]
		if alreadyExists {
			queryParams.params[index].values = append(queryParams.params[index].values, value)
		} else {
			queryParams.params = append(queryParams.params, queryParam{key, []string{value}})
			queryParams.indexes[key] = i
		}

	}

	return queryParams
}
