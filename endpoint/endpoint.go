package endpoint

import (
	"learn_rabbitmq/service/ConsumeService"
	"learn_rabbitmq/service/PublishService"
	"net/http"
)

func RabbitPublishEndpoint(res http.ResponseWriter, req *http.Request) {

	status, err := PublishService.PublishService(req)
	contentRespnse := Response{
		StatusCode: status,
		Content:    nil,
		Success:    true,
	}
	if status != http.StatusOK || err != nil {
		contentRespnse.Success = false
		contentRespnse.Content = err.Error()
	}
	Serve(res, contentRespnse)
}

func RabbitConsumeEndpoint(res http.ResponseWriter, req *http.Request) {

	status, err := ConsumeService.ConsumeService(req)
	contentRespnse := Response{
		StatusCode: status,
		Content:    nil,
		Success:    true,
	}
	if status != http.StatusOK || err != nil {
		contentRespnse.Success = false
		contentRespnse.Content = err.Error()
	}
	Serve(res, contentRespnse)
}
