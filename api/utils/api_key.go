package utils

import "time"

type APIKey struct {
	KeyString  string     `bson:"keyString" json:"keyString"`
	Expiration *time.Time `bson:"expiration,omitempty" json:"expiration,omitempty"`
}

func (apikey *APIKey) IsExpired() bool {
	if apikey.Expiration == nil {
		return false
	}
	return apikey.Expiration.Before(time.Now())
}
