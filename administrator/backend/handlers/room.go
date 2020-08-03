package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JUNNETWORKS/house_administrator/data"
	"github.com/julienschmidt/httprouter"
)

func GetRooms(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rooms, err := data.GetRooms()
	if err != nil {
		log.Println(err)
	}
	output, err := json.Marshal(&rooms)
	if err != nil {
		log.Println(err)
	}
	w.Write(output)
	return
}

func AddRoom(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return
}
