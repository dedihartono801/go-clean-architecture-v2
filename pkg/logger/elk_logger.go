package elk

import "time"

type ElasticLogger struct {
	Method    string    `json:"method"`
	Action    string    `json:"action"`
	Path      string    `json:"path"`
	UserID    int       `json:"user_id"`
	Query     []string  `json:"query"`
	PostData  string    `json:"post_data"`
	IP        string    `json:"ip"`
	Timestamp time.Time `json:"@timestamp"`
	IpAddress string    `json:"ip_address"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	AppId     string    `json:"app_id"`
}
