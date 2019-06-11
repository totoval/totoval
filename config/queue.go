package config

import (
	. "github.com/totoval/framework/config"
	"totoval/app/models"
)

func init() {
	queue := make(map[string]interface{})

	queue["default"] = Env("QUEUE_CONNECTION", "memory")
	queue["connections"] = map[string]interface{}{
		"nsq": map[string]interface{}{
			"nsqd": []map[string]interface{}{
				{
					"tcp": map[string]interface{}{
						"host": Env("QUEUE_NSQD_TCP_HOST", "127.0.0.1"),
						"port": Env("QUEUE_NSQD_TCP_PORT", "4150"),
					},
				},
			},
			"nsqlookupd": []map[string]interface{}{
				{
					"http": map[string]interface{}{
						"host": Env("QUEUE_NSQLOOKUPD_HTTP_HOST", "http://127.0.0.1"),
						"port": Env("QUEUE_NSQLOOKUPD_HTTP_PORT", "4161"),
					},
				},
			},
		},
	}

	queue["max_in_flight"] = Env("QUEUE_MAX_IN_FLIGHT", 25)
	queue["failed_db_processor_model"] = &models.FailedQueue{}

	Add("queue", queue)
}
