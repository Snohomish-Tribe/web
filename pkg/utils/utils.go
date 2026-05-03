package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/snohomishtribe/pkg/models"
)

func FetchEvents(url string, token string) models.EventsPageData {
	godotenv.Load()
	var strapiRes models.EventsPageData

	client := &http.Client{Timeout: time.Second * 10}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("There was no response from Strapi")
			}
		}()
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Unable to read response")
	} else {
		json.Unmarshal(body, &strapiRes)
		fmt.Println(strapiRes)
	}

	defer res.Body.Close()

	return strapiRes
}
