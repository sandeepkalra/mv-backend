package main

import (
	"database/sql"
	"mv/utils"
	"time"
)

type Item struct {
	ID                         int64     `json:"id"`
	Name                       string    `json:"name"`
	AliasName                  string    `json:"alias_name, omitempty"`
	Manufacturer               string    `json:"manufacturer"`
	Owner                      string    `json:"owner, omitempty"`
	CreatedOn                  time.Time `json:"created_on, omitempty"`
	ExpiredOn                  time.Time `json:"expired_on, omitempty"`
	IsExpired                  bool      `json:"is_expired, omitempty"`
	Category                   string    `json:"category"`
	SubCategory                string    `json:"sub_category"`
	SubSubCategory             string    `json:"sub_sub_category, omitempty"`
	SubSubSubCategory          string    `json:"sub_sub_sub_category, omitempty"`
	RegionCountry              string    `json:"region_country, omitempty"`
	RegionState                string    `json:"region_state, omitempty"`
	RegionCity                 string    `json:"region_city, omitempty"`
	RegionPin                  string    `json:"region_pin, omitempty"`
	ItemUrl                    string    `json:"item_url, omitempty"`
	Relationship               string    `json:"relationship, omitempty"`
	RelationshipDurationInDays int       `json:"relationship_duration_in_days, omitempty"`
}

type ItemRequest struct {
	ItemRequested Item   `json:"item"`
	CookieString  string `json:"cookie"`
}

type ItemModule struct {
	DataBase *sql.DB
	RedisDB  *utils.RedisDb
}
