package httpapi

import (
	"encoding/json"
	"net/http"
)

type EventBody struct {
	Id      string                 `json:"event_id"`
	Type    string                 `json:"event_type"`
	Source  string                 `json:"source"`
	Payload map[string]interface{} `json:"payload"`
}

func CreateEvent(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var body EventBody
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		res.WriteHeader(400)
		res.Write([]byte(err.Error()))
	}

	if body.Id == "" {
		res.WriteHeader(400)
		res.Write([]byte("Id is required"))
	}
	if body.Type == "" {
		res.WriteHeader(400)
		res.Write([]byte("Type is required"))
	}
	if body.Source == "" {
		res.WriteHeader(400)
		res.Write([]byte("Source is required"))
	}
	if body.Payload == nil {
		res.WriteHeader(400)
		res.Write([]byte("Payload is required"))
	}
	if len(body.Payload) == 0 {
		res.WriteHeader(400)
		res.Write([]byte("Payload is empty"))
	}
}
