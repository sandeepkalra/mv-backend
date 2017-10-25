
package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SRespJson struct {
	Code     int         `json:"code"`
	Msg string `json:"message"`
	Response interface{} `json:"response"`
}

func GetResponseObject() *SRespJson {
	return &SRespJson{Code: -1, Msg: "Invalid Params", Response: "None"}
}

func (out *SRespJson) Send(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data, err := json.Marshal(*out)
	if err != nil {
		fmt.Println("error marshalling the results", data)
	} else {
		fmt.Fprintf(res, string(data))
	}
	return
}
