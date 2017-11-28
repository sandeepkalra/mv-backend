package main

import (
	"../utils"
	"encoding/json"
	"net/http"

	"../models"
	"strconv"

	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// LookupItem lookup for item, or list of items in a category
func (im *ItemModule) LookupItem(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := ItemRequest{ItemRequested: ItemObj{}, CookieString: ""}
	out := utils.GetResponseObject()
	defer out.Send(res)

	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	/* I choose to do this:
		An ItemObj can be narrowed based on this :
		(a) Name			iPhone
	 	(b) Manufacturer	Apple
		(c) Category		Electronics
		(d) SubCategory 	Phone
		For Simplistic search, I need (a)+(b), (a)+(c)+(d) , (b)+(c)+(d), or all.
	*/

	a := (len(request.ItemRequested.Name) != 0)
	b := (len(request.ItemRequested.Manufacturer) != 0)
	c := (len(request.ItemRequested.Category) != 0)
	d := (len(request.ItemRequested.SubCategory) != 0)

	case1 := a && b
	fmt.Println("case1", case1)
	case2 := a && c && d
	fmt.Println("case2", case2)
	case3 := b && c && d
	fmt.Println("case3", case3)
	case4 := a && b && c && d
	fmt.Println("case4", case4)

	if !(case1 || case2 || case3) {
		out.Msg = "Not enough info to locate the ItemObj(s). "
		return
	}

	var items models.ItemSlice
	var err error

	if case4 {
		items, err = models.Items(im.DataBase,
			qm.Where("name = ? AND manufacturer = ? AND category = ? AND sub_category = ?",
				request.ItemRequested.Name,
				request.ItemRequested.Manufacturer,
				request.ItemRequested.Category,
				request.ItemRequested.SubCategory)).All()
		fmt.Println("case4, len=", len(items))
		if err != nil || len(items) == 0 {
			if err != nil {
				out.Msg = err.Error()
			} else {
				out.Msg = "no items found"
				out.Response = 0
				out.Code = 0
			}
			return
		}

	} else if case1 {
		items, err = models.Items(im.DataBase,
			qm.Where("name = ? AND manufacturer = ? ",
				request.ItemRequested.Name,
				request.ItemRequested.Manufacturer)).All()
		fmt.Println("case1, len=", len(items))
		if err != nil || len(items) == 0 {
			if err != nil {
				out.Msg = err.Error()
			} else {
				out.Msg = "no items found"
				out.Response = 0
				out.Code = 0
			}
			return
		}

	} else if case2 {
		items, err = models.Items(im.DataBase,
			qm.Where("name = ?  AND category = ? AND sub_category = ?",
				request.ItemRequested.Name,
				request.ItemRequested.Category,
				request.ItemRequested.SubCategory)).All()
		fmt.Println("case2, len=", len(items))
		if err != nil || len(items) == 0 {
			if err != nil {
				out.Msg = err.Error()
			} else {
				out.Msg = "no items found"
				out.Response = 0
				out.Code = 0
			}
			return
		}

	} else if case3 {
		items, err = models.Items(im.DataBase,
			qm.Where("manufacturer = ? AND category = ? AND sub_category = ?",
				request.ItemRequested.Manufacturer,
				request.ItemRequested.Category,
				request.ItemRequested.SubCategory)).All()

		fmt.Println("case3, len=", len(items))
		if err != nil || len(items) == 0 {
			if err != nil {
				out.Msg = err.Error()
			} else {
				out.Msg = "no items found"
				out.Response = 0
				out.Code = 0
			}
			return
		}
	}

	/* At this point, we have a narrowed list of items(ItemSlice) */
	itemSliceResp := make([]ItemObj, 0, 10)
	for _, i := range items {

		item := ItemObj{
			ID:                i.ID,
			Name:              i.Name.String,
			Manufacturer:      i.Manufacturer.String,
			Category:          i.Category.String,
			SubCategory:       i.SubCategory.String,
			SubSubCategory:    i.SubSubCategory.String,
			SubSubSubCategory: i.SubSubSubCategory.String,
			RegionCountry:     i.RegionCountry.String,
			RegionState:       i.RegionCity.String,
			RegionCity:        i.RegionCity.String,
			RegionPin:         i.RegionPin.String,
			AliasName:         i.AliasName.String,
			ItemURL:           i.Itemurl.String,
			Owner:             i.OwnerName.String,
			CreatedOn:         &i.CreatedOn.Time,
			ExpiredOn:         &i.ExpiryOn.Time,
			IsExpired:         i.HasExpired.Valid,
		}

		itemSliceResp = append(itemSliceResp, item)
	}

	out.Code = 0
	out.Msg = strconv.Itoa(len(itemSliceResp)) + " items found."
	out.Response = map[string]interface{}{
		"items": itemSliceResp,
	}
	return
}
