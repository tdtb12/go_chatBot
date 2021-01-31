package controller

import (
	"bytes"
	"chatBot/constants"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RequestBody struct {
	Destination string   `json:"destination"`
	Events      []Events `json:"events"`
}

type Events struct {
	ReplyToken string  `json:"replyToken"`
	Type       string  `json:"type"`
	Mode       string  `json:"mode"`
	Timestamp  int64   `json:"timestamp"`
	Source     Source  `json:"source"`
	Message    Message `json:"message"`
}

type Source struct {
	Type   string `json:"type"`
	UserId string `json:"userId"`
}

type Message struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

func RootController(w http.ResponseWriter, r *http.Request) {
	log.Println("RootController " + constants.APIStart)

	header := r.Header.Get("X-Line-Signature")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	decoder := json.NewDecoder(r.Body)
	var req RequestBody
	err = decoder.Decode(&req)
	if err != nil {
		// log error and return 400 to caller
		log.Fatal(err)
	}

	decoded, err := base64.StdEncoding.DecodeString(header)
	if err != nil {
		// ...
	}
	hash := hmac.New(sha256.New, []byte("ea4aa0e4aa9bd5c30ae856345989d8aa"))
	hash.Write(body)

	if hmac.Equal(hash.Sum(nil), decoded) {
		fmt.Println("equal")
		fmt.Println(ioutil.NopCloser(bytes.NewBuffer(body)))
		bodyString := string(body)
		fmt.Println(bodyString)
	} else {
		log.Fatal("not equal")
	}
	log.Println("RootController " + constants.APIEnd)
}
