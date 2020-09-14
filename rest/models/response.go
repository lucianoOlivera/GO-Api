package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status     int         `json:"status"`
	Data       interface{} `json:"data"`
	Messenge   string      `json:"messenge"`
	contenType string
	writer     http.ResponseWriter
}

func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{Status: http.StatusOK, writer: w, contenType: "application/json"}
}

func SendNotFoud(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFoud()
	response.Send()

}

func SendUnProcessableEntity(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.UnProcessableEntity()
	response.Send()
}
func (this *Response) UnProcessableEntity() {
	this.Status = http.StatusUnprocessableEntity
	this.Messenge = "UnProcessableEntity"

}
func (this *Response) NotFoud() {
	this.Status = http.StatusNotFound
	this.Messenge = "Not found"

}
func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}

func (this *Response) Send() {
	this.writer.Header().Set("Content-Type", this.contenType)
	this.writer.WriteHeader(this.Status)

	output, _ := json.Marshal(&this)
	fmt.Fprintf(this.writer, string(output))

}
func SendNoContent(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotContent()
	response.Send()
}
func (this *Response) NotContent() {
	this.Status = http.StatusNoContent
	this.Messenge = "no conetnt"

}
