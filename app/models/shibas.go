package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiUrl = "http://shibe.online/api/shibes?"

type ShibaImgs []string

func GetShibas(count string, urls string, httpsUrls string) (imgurls ShibaImgs) {

	url := fmt.Sprintf("%scount=%s&urls=%s&httpsUrls=%s", apiUrl, count, urls, httpsUrls)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	var res ShibaImgs

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println(err)
	}
	return res
}
