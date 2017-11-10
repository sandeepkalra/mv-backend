package main

import (
	"encoding/json"
	"mv/models"
	"mv/utils"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
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

	if request.ManufacturerContains == "All" {
		request.ManufacturerContains = ""
	}

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

	categories := make([]CategoryList, 1)

	if request.NeedCategoryList {
		cList, e := models.Categories(im.DataBase, qm.Where("fk_parent_category_id = ?", 0)).All()
		if e == nil {
			category := CategoryList{}
			categoryID := int64(0)
			for _, m := range cList {
				if len(request.CategoryNameContains) > 0 { /* filter */
					if strings.ContainsAny(m.Name.String, request.ManufacturerContains) == true {
						category.CategoryName = m.Name.String
						categoryID = m.ID
					}
				} else {
					category.CategoryName = m.Name.String
					categoryID = m.ID
				}
				if request.NeedSubCategoryList {
					listSubCategories, err := models.Categories(im.DataBase, qm.Where("fk_parent_category_id=?", categoryID)).All()
					if err == nil {
						for _, m := range listSubCategories {
							if len(request.CategoryNameContains) > 0 { /* filter */
								if strings.ContainsAny(m.Name.String, request.CategoryNameContains) {
									category.SubCategoryNames = append(category.SubCategoryNames, m.Name.String)
								}
							} else {
								category.SubCategoryNames = append(category.SubCategoryNames, m.Name.String)
							}
						}
					}
				}
				categories = append(categories, category)
			}
		}
		response["categories"] = categories
	}

	out.Code = 0
	out.Msg = "ok"
	out.Response = response

	return

}
