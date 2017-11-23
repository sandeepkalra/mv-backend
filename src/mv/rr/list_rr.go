package main

import (
	"../models"
	"../utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// ListRR list reviews and ratings
func (rr *RRModule) ListRR(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	request := RR{}
	out := utils.GetResponseObject()
	defer out.Send(res)
	if e := json.NewDecoder(req.Body).Decode(&request); e != nil {
		out.Msg = " failed to decode incoming msg "
		return
	}

	if request.ItemID == 0 {
		out.Msg = "requested item id cannot be zero"
		return
	}

	reviews, err := models.ReviewRatingRelationships(rr.DataBase, qm.Where("fk_item_id =?", request.ItemID)).All()
	if err != nil || len(reviews) == 0 {
		out.Msg = "could not find the required item, or zero reviews found"
		return
	}

	responses := make([]RR, 10)

	cnt := 0
	for _, r := range reviews {
		cnt++
		resp := RR{
			PersonID:         r.FKPersonID,
			ItemID:           r.FKItemID,
			Relationship:     r.MyRelationshipWithItem.String,
			Comments:         r.Comments.String,
			Rating:           r.Rating.Int,
			Pros:             r.Pros.String,
			Cons:             r.Cons.String,
			RelationshipDate: r.Createdon.Time,
			HideDetails:      r.HideDetails.Valid,
			IsResponse:       r.IsResponse.Valid,
			HasResponse:      r.HasResponse.Valid,
		}

		if r.HideDetails.Valid {
			resp.PersonID = 0
			resp.Relationship = " hidden on request from user. "
		} else {
			fullName := ""
			p, e := models.FindPerson(rr.DataBase, resp.PersonID)
			if p != nil && e == nil {
				fullName = p.LName.String + ", " + p.FName.String
			}
			resp.PersonName = fullName
		}

		responses = append(responses, resp)

		/* we fetch all responses, and add them immidiately after */
		if resp.HasResponse {
			respToResp, e := models.ReviewRatingRelationships(rr.DataBase, qm.Where("fk_item_id = ?", resp.ItemID)).All()
			if e == nil && len(respToResp) != 0 {
				for _, rToR := range respToResp {
					cnt++
					rtr := RR{
						PersonID:         rToR.FKPersonID,
						ItemID:           rToR.FKItemID,
						Relationship:     rToR.MyRelationshipWithItem.String,
						Comments:         rToR.Comments.String,
						Rating:           rToR.Rating.Int,
						Pros:             rToR.Pros.String,
						Cons:             rToR.Cons.String,
						RelationshipDate: rToR.Createdon.Time,
						HideDetails:      rToR.HideDetails.Valid,
						IsResponse:       rToR.IsResponse.Valid,
						HasResponse:      rToR.HasResponse.Valid,
					}
					responses = append(responses, rtr) /* add response to response to the list */
				}
			} // no error fetching respToResp slice
		}
	}

	out.Code = 0
	out.Msg = strconv.Itoa(cnt) + " reviews found."
	out.Response = map[string]interface{}{
		"reviews": responses,
	}
	return
}
