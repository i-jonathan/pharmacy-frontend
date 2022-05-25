package global

import (
	"PharmUI/models"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func Getter(endpoint string, respStruct interface{}) error {
	baseURL := ApiURL
	fullURL := baseURL + endpoint

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

	err = json.NewDecoder(resp.Body).Decode(respStruct)
	if err != nil {
		log.Println("JSON Decoder error: ", err)
		return err
	}
	return nil
}

func Poster(endpoint string, requestBody interface{}) (models.Response, error) {
	fullUrl := ApiURL + endpoint

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(requestBody)
	if err != nil {
		log.Println(err)
		return models.Response{}, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, fullUrl, b)

	if err != nil {
		log.Println(err)
		return models.Response{}, err
	}

	resp, err := client.Do(req)

	var respStruct models.Response
	err = json.NewDecoder(resp.Body).Decode(&respStruct)
	if err != nil {
		log.Println(err)
		return models.Response{}, err
	}
	respStruct.Status = resp.StatusCode
	return respStruct, err
}