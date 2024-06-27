package accessgen

import (
	"fmt"
	"io"
	"net/http"
)

func AccessGen(client_id, client_secret string) (string, error) {
	urll := fmt.Sprintf("https://hh.ru/oauth/token?grant_type=client_credentials&client_id=%s&client_secret=%s", client_id, client_secret)
	res, err := http.Post(urll, "text/plain", nil)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
