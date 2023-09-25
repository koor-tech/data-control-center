package httpapi

var PingResponse = Ping{
	Message: "Pong",
}

type Ping struct {
	Message string `json:"message"`
}
