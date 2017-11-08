package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"mv/models"
	"mv/utils"
	"net/http"
	"strings"
)

//LookupList returns list of manufacturer and categories, and many more things to come.
func (im *ItemModule) LookupList(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := GetListRequest{CookieString: ""}
	out := utils.GetResponseObject()
	defer out.Send(res)

	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	if request.NeedManufacturerList == false &&
		request.NeedCategoryList == false {
		out.Msg = "can't makeout what you want! "
		return
	}
	response := make(map[string]interface{})

	manufacturers := make([]string, 1)

	if request.NeedManufacturerList {
		mList, e := models.ManufacturersLists(im.DataBase).All()
		if e != nil {
			for _, m := range mList {
				if len(request.ManufacturerContains) > 0 {
					if strings.ContainsAny(m.Name.String, request.ManufacturerContains) == true {
						manufacturers = append(manufacturers, m.Name.String)
					}
				} else {
					manufacturers = append(manufacturers, m.Name.String)
				}
			}
		}
		response["manufacturers"] = manufacturers
	}

	categories := make([]string, 1)
	if request.NeedCategoryList {
		/* We lookup only top level categories */
		cList, e := models.Categories(im.DataBase, qm.Where("fk_parent_category_id = ?", 0)).All()
		if e != nil {
			for _, m := range cList {
				if len(request.CategoryNameContains) > 0 {
					if strings.ContainsAny(m.Name.String, request.ManufacturerContains) == true {
						categories = append(categories, m.Name.String)
					}
				} else {
					categories = append(categories, m.Name.String)
				}
			}
		}
		response["categories"] = categories
	}

	out.Code = 0
	out.Msg = "ok"
	out.Response = response

	return

}
