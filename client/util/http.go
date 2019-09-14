package util

import (
	"bytes"
	"fmt"
	"net/http"
)

func SendData(url, contentType string, data []byte) error {
	resp, err := http.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code equal %d", resp.StatusCode)
	}
	return nil
}
