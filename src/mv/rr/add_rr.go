package main

import (
	"encoding/json"
	"mv/models"
	"mv/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
	null "gopkg.in/volatiletech/null.v6"
)

func (rr *RRModule) AddRR(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := RR{}
	out := utils.GetResponseObject()
	defer out.Send(res)
	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	if request.PersonId == 0 ||
		request.ItemId == 0 ||
		len(request.Comments) == 0 ||
		request.Rating > 10 ||
		request.Rating < 10 ||
		len(request.Relationship) == 0 {
		out.Msg = "invalid rating, or, zero comments, or, invalid Ids "
		return
	}

	review := models.ReviewRatingRelationship{
		FKPersonID:             request.PersonId,
		FKItemID:               request.ItemId,
		MyRelationshipWithItem: null.StringFrom(request.Relationship),
		Comments:               null.StringFrom(request.Comments),
		Rating:                 null.IntFrom(request.Rating),
		Pros:                   null.StringFrom(request.Pros),
		Cons:                   null.StringFrom(request.Cons),
		Createdon:              null.TimeFrom(request.RelationshipDate),
	}

	if err := review.Insert(rr.DataBase); err != nil {
		out.Msg = err.Error()
		return
	}

	out.Code = 0
	out.Msg = "ok"
	out.Response = map[string]interface{}{
		"review_id": review.ID,
	}
	return

}
