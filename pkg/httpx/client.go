package httpx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post(url string, body interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http response code: %d", resp.StatusCode)
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := new(response)
	if err := json.Unmarshal(b, res); err != nil {
		return err
	}

	if res.Code != 0 {
		return errors.New(res.Message)
	}

	return nil
}
