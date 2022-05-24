package global

import (
	"encoding/json"
	"log"
	"net/http"
)

func Getter(endpoint string, respStruct interface{}) error {
	baseURL := ApiURL
	fullURL := baseURL + endpoint

	result := respStruct
	client := &http.Client{}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		log.Println("Can't create new request: ", err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error with getting from"+endpoint+" : ", err)
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		log.Println("JSON Decoder error: ", err)
		return err
	}
	return nil
}
