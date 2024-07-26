package endpoint

import (
	"encoding/json"
	"learn_rabbitmq/model/errorModel"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Content    interface{} `json:"content"`
	Success    bool        `json:"success"`
}

func Serve(res http.ResponseWriter, content Response) {

	b, _ := json.Marshal(&content)

	res.WriteHeader(content.StatusCode)
	res.Header().Add("Content-Type", "application/json")
	_, err := res.Write(b)
	if err != nil {
		errorModel.FailOnErr(err, "Failed write response...")
	}
}
