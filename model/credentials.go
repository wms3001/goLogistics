package model

type Credentials struct {
	Base
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}
