package api

import (
	"github.com/gorilla/mux"
	"github.com/mathieupassenaud/leds/backend"
	"net/http"
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
	leds []struct {
		index int `json:"index"`
		red   int `json:"red"`
		green int `json:"green"`
		blue  int `json:"blue"`
	}
}

func (handler *HttpSetHandler) HandlePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	input := inputJson{}
	json.Unmarshal(body, &input)
	for i := 0; i < len(input.leds); i = i + 1 {
		s := statuses.ChangeStatus(input.leds[i].index, statuses.Color{input.leds[i].red, input.leds[i].green, input.leds[i].blue}, 1, 0)
		renderer.Apply(s)
	}
	renderer.ForceRender()
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
