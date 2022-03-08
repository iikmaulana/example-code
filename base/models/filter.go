package models

import "github.com/uzzeet/uzzeet-gateway/models"

type FilterParams struct {
	Page     int64
	Limit    int64
	SortType string
	Filter   []string
	Search   string
	All      string
	Oauth    models.AuthorizationInfo
}

type FormMetaData struct {
	TotalDataPerpage int
	TotalPage        int
	TotalData        int64
	To               int
	From             int
}
