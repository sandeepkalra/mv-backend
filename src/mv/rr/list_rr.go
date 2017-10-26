package main

import (
	"encoding/json"
	"mv/models"
	"mv/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (rr *RRModule) ListRR(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := RR{}
	out := utils.GetResponseObject()
	defer out.Send(res)
	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	if request.ItemId == 0 {
		out.Msg = "requested item id cannot be zero"
		return
	}

	reviews, err := models.ReviewRatingRelationships(rr.DataBase, qm.Where("fk_item_id =?", request.ItemId)).All()
	if err != nil || len(reviews) == 0 {
		out.Msg = "could not find the required item, or zero reviews found"
		return
	}

	responses := make([]RR, 10)

	cnt := 0
	for _, r := range reviews {
		cnt++
		resp := RR{
			PersonId:         r.FKPersonID,
			ItemId:           r.FKItemID,
			Relationship:     r.MyRelationshipWithItem.String,
			Comments:         r.Comments.String,
			Rating:           r.Rating.Int,
			Pros:             r.Pros.String,
			Cons:             r.Cons.String,
			RelationshipDate: r.Createdon.Time,
			HideDetails:      r.HideDetails.Valid,
		}
		if r.HideDetails.Valid {
			resp.PersonId = 0
			resp.Relationship = " hidden on request from user. "
		} else {
			full_name := ""
			p, e := models.FindPerson(rr.DataBase, resp.PersonId)
			if p != nil && e == nil {
				full_name = p.LName.String + ", " + p.FName.String
			}
			resp.PersonName = full_name
		}
		responses = append(responses, resp)
	}
	out.Code = 0
	out.Msg = strconv.Itoa(cnt) + " reviews found."
	out.Response = map[string]interface{}{
		"reviews": responses,
	}
	return
}
