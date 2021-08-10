package keyspot

import (
	"net/http"
	"encoding/json"
	"time"
	"io/ioutil"
	"os"
	"bytes"
	"errors"
	"strconv"
)

var url string = "https://database-driver-ifhogzjzbq-uc.a.run.app"

func GetRecord(accessKey string) (map[string]string, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Get(url + "/" + accessKey)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var record map[string]string

	err = json.Unmarshal(body, &record)

	if err != nil {
		return nil, err
	}

	return record, nil
}

func SetEnvironment(accessKey string) error {
	record, err := GetRecord(accessKey)

	if err != nil {
		return err
	}

	for key, value := range record {
		os.Setenv(key, value)
	}

	return nil
}

func UpdateRecord(accessKey string, newRecord map[string]string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	payload, err := json.Marshal(newRecord)

	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, url + "/" + accessKey, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Response status: " + strconv.Itoa(resp.StatusCode))
	}

	return nil
}