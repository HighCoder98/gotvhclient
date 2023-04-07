package gotvhclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type TvhClient struct {
	Username string
	Password string
	Address  string
}

func NewTvhClient(address string, username string, password string) (*TvhClient, error) {
	client := &TvhClient{
		Username: username,
		Password: password,
		Address:  address,
	}

	return client, nil
}

func (client *TvhClient) apiGetRequest(endpoint string) (string, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", client.Address+"/"+endpoint, nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(client.Username, client.Password)
	response, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {
		panic("Unexpected status code: " + response.Status)
	}

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	return string(b), nil
}

func (client *TvhClient) GetChannels() ([]Channel, error) {
	resp, err := client.apiGetRequest("channel/grid")
	if err != nil {
		return nil, err
	}

	var channelResp ChannelResponse
	err = json.Unmarshal([]byte(resp), &channelResp)

	if err != nil {
		panic(err)
	}

	return channelResp.Entries, nil
}
