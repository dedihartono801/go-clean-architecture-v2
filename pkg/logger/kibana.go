package elk

import (
	"encoding/json"
	"os"

	http_client "github.com/dedihartono801/go-clean-architecture-v2/pkg/helpers"
)

func Kibana(el ElasticLogger) error {
	postDataByte, _ := json.Marshal(el)

	_, err := http_client.PostHTTPRequest(os.Getenv("ELASTIC_LOG_URL"), "", postDataByte)
	if err != nil {
		return err
	}

	return nil
}
