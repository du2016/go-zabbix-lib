package zabbix

import (
	"github.com/du2016/go-zabbix-lib/reflector"
	"log"
)

type (
	Eventsource string
	Status      string
)

const (
	Enable  Status = "0"
	Disable Status = "1"
)
const (
	Triggeraction   Eventsource = "0"
	Discoveryaction Eventsource = "1"
	Autoregaction   Eventsource = "2"
	Internalaction  Eventsource = "3"
)

type Action struct {
	Actionid          string      `json:"actionid,omitempty"`
	Name              string      `json:"name"`
	ActionEventsource Eventsource `json:"eventsource"` //0trigger，1 discovery 2auto registration" 3 internal`
	ActionStatus      Status      `json:"status"`
}

type Actions []Action

func (api *API) ActionGet(params Params) (res Actions, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	response, err := api.CallWithError("action.get", params)
	if err != nil {
		return
	}
	// fmt.Printf("$v%", response.Result, "111")
	reflector.MapsToStructs2(response.Result.([]interface{}), &res, reflector.Strconv, "json")
	return
}

func (api *API) ActionGetById(id string) (res *Action, err error) {
	actions, err := api.ActionGet(Params{"filter": map[string]string{"actionid": id}})
	if err != nil {
		return
	}

	if len(actions) == 1 {
		res = &actions[0]
	} else {
		e := ExpectedOneResult(len(actions))
		err = &e
	}
	return
}

func (api *API) ActionGetByStatus(s Status, e Eventsource) (res Actions, err error) {
	return api.ActionGet(Params{"filter": map[string]string{"status": string(s), "eventsource": string(e)}})
}

func (api *API) ActionUpdatestatusByStatus(s Status, e Eventsource, d Status) {
	var res Actions
	var err error
	if s == "" {
		res, err = api.ActionGet(Params{"filter": map[string]string{"eventsource": string(e)}})
		if err != nil {
			log.Fatal(err)
		}
	}
	res, err = api.ActionGet(Params{"filter": map[string]string{"status": string(s), "eventsource": string(e)}})
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range res {
		if v.ActionStatus != d {
			api.ActionUpdateByid(v.Actionid, d)
		}
	}
}

func (api *API) ActionUpdateByid(id string, s Status) error {
	_, err := api.CallWithError("action.update", Params{"actionid": id, "status": s})
	return err
}
