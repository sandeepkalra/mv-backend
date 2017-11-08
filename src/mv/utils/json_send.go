package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SRespJSON is the protocol of backend
type SRespJSON struct {
	Code     int         `json:"code"`
	Msg      string      `json:"message"`
	Response interface{} `json:"response"`
}

// GetResponseObject constructs and sends a default obj
func GetResponseObject() *SRespJSON {
	return &SRespJSON{Code: -1, Msg: "Invalid Params", Response: "None"}
}

// Send marshal and send response out using responseWriter
func (out *SRespJSON) Send(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data, err := json.Marshal(*out)
	if err != nil {
		fmt.Println("error marshalling the results", data)
	} else {
		fmt.Fprintf(res, string(data))
	}
	return
}
