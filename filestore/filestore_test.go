package filestore

import "testing"

func TestStoringAndGetting(t *testing.T) {

	FilestoreURL = "http://localhost:10000/"

	key := "testing"
	data := "Hello World"

	body := []byte(data)

	err := PutItem(key, body, "")
	if err != nil {
		t.Errorf("Failed to put item. Error: %v", err)
	}

	resp, err := GetItem(key)
	if err != nil {
		t.Errorf("Failed to get item. Error: %v", err)
	}

	if string(resp) != data {
		t.Errorf("Sent and recieved does not equal. %v vs %v", data, resp)
	}
}
