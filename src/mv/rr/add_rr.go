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
		HasResponse:            null.Int8From(0), /* we overwrite later if true */
		IsResponse:             null.Int8From(0), /* we overwrite later if true */
	}
	if request.HasResponse { /* Overwrite , we need this step because golang do not allow bool to be converted to int. */
		review.HasResponse = null.Int8From(1)
	}
	if request.IsResponse { /* Overwrite , we need this step because golang do not allow bool to be converted to int. */
		review.IsResponse = null.Int8From(1)
		/* if this review is a response to other review, then we have to update both */
		orig_review, e := models.FindReviewRatingRelationship(rr.DataBase, request.ItemId)
		if e == nil && orig_review != nil {
			orig_review.HasResponse = null.Int8From(1)
			orig_review.Update(rr.DataBase)
		}
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
