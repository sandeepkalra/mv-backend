package main

import (
	"encoding/json"
	"mv/models"
	"mv/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"gopkg.in/volatiletech/null.v6"
)

// AddItem add item to database
func (im *ItemModule) AddItem(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := ItemRequest{ItemRequested: Item{}, CookieString: ""}
	out := utils.GetResponseObject()
	defer out.Send(res)

	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	if len(request.ItemRequested.Name) == 0 ||
		len(request.ItemRequested.Manufacturer) == 0 ||
		len(request.ItemRequested.Category) == 0 ||
		len(request.ItemRequested.SubCategory) == 0 {
		out.Msg = "name, manufacturer/creator and category/sub_category cannot be zero."
		return
	}

	_item, e := models.Items(im.DataBase, qm.Where("name = ? AND manufacturer = ? AND category = ? AND sub_category = ?",
		request.ItemRequested.Name,
		request.ItemRequested.Manufacturer,
		request.ItemRequested.Category,
		request.ItemRequested.SubCategory)).
		One()

	if e == nil && _item != nil {
		out.Msg = "Item exist in our database"
		out.Response = map[string]interface{}{
			"item": Item{
				ID:                _item.ID,
				Name:              _item.Name.String,
				Manufacturer:      _item.Manufacturer.String,
				Category:          _item.Category.String,
				SubCategory:       _item.SubCategory.String,
				SubSubCategory:    _item.SubSubCategory.String,
				SubSubSubCategory: _item.SubSubSubCategory.String,
				RegionCountry:     _item.RegionCountry.String,
				RegionState:       _item.RegionCity.String,
				RegionCity:        _item.RegionCity.String,
				RegionPin:         _item.RegionPin.String,
				AliasName:         _item.AliasName.String,
				ItemURL:           _item.Itemurl.String,
				Owner:             _item.OwnerName.String,
				CreatedOn:         _item.CreatedOn.Time,
				ExpiredOn:         _item.ExpiryOn.Time,
				IsExpired:         _item.HasExpired.Valid,
			},
		}
		return
	}

	/* Else, we are here to add this item that user has requested */
	item := models.Item{
		ID:                request.ItemRequested.ID,
		Name:              null.StringFrom(request.ItemRequested.Name),
		Manufacturer:      null.StringFrom(request.ItemRequested.Manufacturer),
		Category:          null.StringFrom(request.ItemRequested.Category),
		SubCategory:       null.StringFrom(request.ItemRequested.SubCategory),
		SubSubCategory:    null.StringFrom(request.ItemRequested.SubSubCategory),
		SubSubSubCategory: null.StringFrom(request.ItemRequested.SubSubSubCategory),
		RegionCountry:     null.StringFrom(request.ItemRequested.RegionCountry),
		RegionState:       null.StringFrom(request.ItemRequested.RegionCity),
		RegionCity:        null.StringFrom(request.ItemRequested.RegionCity),
		RegionPin:         null.StringFrom(request.ItemRequested.RegionPin),
		AliasName:         null.StringFrom(request.ItemRequested.AliasName),
		Itemurl:           null.StringFrom(request.ItemRequested.ItemURL),
		OwnerName:         null.StringFrom(request.ItemRequested.Owner),
		CreatedOn:         null.TimeFrom(request.ItemRequested.CreatedOn),
		ExpiryOn:          null.TimeFrom(request.ItemRequested.ExpiredOn),
	}

	if request.ItemRequested.IsExpired {
		item.HasExpired = null.Int8From(1)
	} else {
		item.HasExpired = null.Int8From(0)
	}

	if e := item.Insert(im.DataBase); e != nil {
		out.Msg = e.Error()
		return
	}

	/* At this point, we lookup manufacturer_list, and categories, and if they need to have this, we add there too */

	if _manufacturer, err := models.ManufacturersLists(im.DataBase, qm.Where("name=?", request.ItemRequested.Manufacturer)).One(); err != nil {
		_manufacturer.Name = null.StringFrom(request.ItemRequested.Manufacturer)
		_manufacturer.Insert(im.DataBase)
	}

	var _category0ID, _category1ID, _category2ID int64

	/* Category, level-0 */
	if _category, err := models.Categories(im.DataBase, qm.Where("name=? AND fk_parent_category_id=?", request.ItemRequested.Category, 0)).One(); err != nil {
		/* fk_parent_category_id = 0 , i.e. no parent of this category */
		_category.Name = null.StringFrom(request.ItemRequested.Category)
		_category.FKParentCategoryID = null.Int64From(0)
		_category.Insert(im.DataBase)
		_category0ID = _category.ID
	} else {
		_category0ID = _category.ID
	}

	/* Category, level - 1 */
	if _category, err := models.Categories(im.DataBase, qm.Where("name=? AND fk_parent_category_id=?", request.ItemRequested.Category, _category0ID)).One(); err != nil {
		/* fk_parent_category_id = 0 , i.e. no parent of this category */
		_category.Name = null.StringFrom(request.ItemRequested.Category)
		_category.FKParentCategoryID = null.Int64From(_category0ID)
		_category.Insert(im.DataBase)
		_category1ID = _category.ID
	} else {
		_category1ID = _category.ID
	}

	/* Category, level - 2 */
	if _category, err := models.Categories(im.DataBase, qm.Where("name=? AND fk_parent_category_id=?", request.ItemRequested.Category, _category1ID)).One(); err != nil {
		/* fk_parent_category_id = 0 , i.e. no parent of this category */
		_category.Name = null.StringFrom(request.ItemRequested.Category)
		_category.FKParentCategoryID = null.Int64From(_category1ID)
		_category.Insert(im.DataBase)
		_category2ID = _category.ID
	} else {
		_category2ID = _category.ID
	}

	/* Category, level - 3 */
	if _category, err := models.Categories(im.DataBase, qm.Where("name=? AND fk_parent_category_id=?", request.ItemRequested.Category, _category2ID)).One(); err != nil {
		/* fk_parent_category_id = 0 , i.e. no parent of this category */
		_category.Name = null.StringFrom(request.ItemRequested.Category)
		_category.FKParentCategoryID = null.Int64From(_category2ID)
		_category.Insert(im.DataBase)
	}

	/* Done , Add success */
	out.Code = 0
	out.Msg = "created"
	out.Response = map[string]interface{}{
		"item_id": item.ID,
	}

	return

}
