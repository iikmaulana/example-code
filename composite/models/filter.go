package models

import "github.com/iikmaulana/example-code/composite/lib"

//todo template filter search
type FilterParams struct {
	Page     int64
	Limit    int64
	SortType string
	Filter   []string
	Search   string
	All      string
	Oauth    lib.AuthorizationInfo
}
