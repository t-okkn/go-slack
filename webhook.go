package goslack

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	_ "fmt"        // for debug
	_ "io/ioutil"  // for debug
)

type SlackMessageToken struct {
	WebhookUrl  string
}

func NewSlackMessageClient(whurl string) *SlackMessageToken {
	return &SlackMessageToken{ whurl }
}

func (t *SlackMessageToken) SendSimpleMessage(msg string) error {
	m := map[string]string{ "text": msg }
	return t.SendMessage(m)
}

func (t *SlackMessageToken) SendMessage(obj interface{}) error {
	p, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	val := url.Values{}
	val.Add("payload", string(p))

	return t.sendPostForm(val)
}

func (t *SlackMessageToken) sendPostForm(payload url.Values) error {
	requri, err := url.ParseRequestURI(t.WebhookUrl)
	if err != nil {
		return err
	}

	resp, err := http.PostForm(requri.String(), payload)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		e := "[status code: " + resp.Status + "] Failed to send data"
		return errors.New(e)
	}

	/* ***********for debug*********** //
	defer resp.Body.Close()
	fmt.Println(resp.Header)
	fmt.Println("*******")
	str, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(str))
	// ******************************* */

	return nil
}

