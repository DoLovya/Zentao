package Api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const TokenUrl = "api.php/v1/tokens"

func GetToken(account string, password string, token *string) error {
	body := map[string]interface{}{
		"account":  account,
		"password": password,
	}
	jsonBody, _ := json.Marshal(&body)
	resp, _ := http.Post(URL+TokenUrl, "application/json", bytes.NewBuffer(jsonBody))
	defer resp.Body.Close()

	response := struct {
		Token string `json:"token"`
	}{}
	if err := UnmarshalResponse(resp, &response); err != nil {
		return err
	}
	*token = response.Token
	return nil
}

func LoadToken(account string, password string) error {
	var token string
	if err := GetToken(account, password, &token); err != nil {
		return err
	}
	if token == "" {
		return fmt.Errorf("禅道账号或密码错误")
	}
	Cache.Token = token
	return nil
}
