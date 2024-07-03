package file

import (
	"encoding/base64"
	"io"
	"net/http"
)

func DownloadFileBase64(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	base64String := base64.StdEncoding.EncodeToString(data)
	return base64String, nil
}
