package lib

var secret_key_super_apps = []byte("um_super_apps")

type AuthorizationInfo struct {
	UserID         string   `json:"id"`
	Username       string   `json:"username"`
	IsOrgAdmin     int      `json:"isorgadmin"`
	IsActive       int      `json:"isactive"`
	OrganizationId string   `json:"organizationid"`
	AppId          string   `json:"app"`
	Exp            int      `json:"exp"`
	UserAccess     []string `json:"user_access"`
}
