package main

import (
	"database/sql"
	"mv/utils"
	"time"
)

// Item item schema to code mapping
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
	ItemURL                    string    `json:"item_url, omitempty"`
	Relationship               string    `json:"relationship, omitempty"`
	RelationshipDurationInDays int       `json:"relationship_duration_in_days, omitempty"`
}

//ItemRequest typical request of item details
type ItemRequest struct {
	ItemRequested Item   `json:"item"`
	CookieString  string `json:"cookie"`
}

// GetListRequest get the list of manufacturer, categories , and many more
type GetListRequest struct {
	NeedManufacturerList bool   `json:"need_manufacturer_list"`
	ManufacturerContains string `json:"manufacturer_name_contains"`

	NeedCategoryList     bool   `json:"need_category_list"`
	CategoryNameContains string `json:"category_name_contains"`

	CookieString string `json:"cookie"`
}

// ItemModule module struct
type ItemModule struct {
	DataBase *sql.DB
	RedisDB  *utils.RedisDb
}
