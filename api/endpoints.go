package api

import (
	"github.com/gorilla/mux"
	"github.com/mathieupassenaud/leds/backend"
	"net/http"
	"io"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func CreateRouter() *mux.Router {

	httpSetHandler := HttpSetHandler{}
	httpAlertHandler := HttpAlertHandler{}
	httpStatusHandler := HttpStatusHandler{}

	r := mux.NewRouter()

	r.HandleFunc("/api/set", (httpSetHandler.HandlePost)).Methods("POST")
	r.HandleFunc("/api/alert", (httpAlertHandler.HandleAlert)).Methods("GET", "POST")
	r.HandleFunc("/api/status", (httpStatusHandler.HandleStatus)).Methods("GET")

	return r
}

type HttpSetHandler struct {
}

func NewHttpSetHandler() *HttpSetHandler {
	return &HttpSetHandler{}
}

type inputJson struct {
	Leds []singleLed `json:"array"`
}

type singleLed struct {
	Index int `json:"index"`
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}


func (handler *HttpSetHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	input := inputJson{}
	json.Unmarshal(body, &input)
	fmt.Println(input);
	for i := 0; i < len(input.Leds); i = i + 1 {
		s := backend.ChangeStatus(input.Leds[i].Index, backend.Color{Red: input.Leds[i].Red, Green: input.Leds[i].Green, Blue: input.Leds[i].Blue}, 1, 0)
		backend.Apply(s)
	}
	backend.ForceRender()
}

type HttpAlertHandler struct {
}

func NewHttpAlertHandler() *HttpAlertHandler {
	return &HttpAlertHandler{}
}

func (handler *HttpAlertHandler) HandleAlert(w http.ResponseWriter, r *http.Request) {

}

type HttpStatusHandler struct {
}

func NewHttpStatusHandler() *HttpStatusHandler {
	return &HttpStatusHandler{}
}

func (handler *HttpStatusHandler) HandleStatus(w http.ResponseWriter, r *http.Request) {

}
