package filestore

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//FilestoreURL is the base URL for the FileStore service
var FilestoreURL string

//DefaultContentType is the default content type sent on all requests if "" is passed as parameter. Defaults to "text/plain"
var DefaultContentType = "text/plain"

//PutItem Posts the data provided to the FileStore service to store under the key provided.
func PutItem(key string, data []byte, contentType string) error {
	body := bytes.NewReader(data)

	if contentType == "" {
		contentType = DefaultContentType
	}

	resp, err := http.Post(FilestoreURL+key, contentType, body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

//GetItem returns the data stored behind the key as a array of bytes
func GetItem(key string) ([]byte, error) {
	resp, err := http.Get(FilestoreURL + key)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
