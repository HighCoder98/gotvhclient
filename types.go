package gotvhclient

type ChannelResponse struct {
	Entries []Channel `json:"entries"`
}

type Channel struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}
