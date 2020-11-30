package http

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"wait4it/model"
)

//HttpChecker ...
type HttpChecker struct {
	Url    string
	Status int
	Text   string
}

//BuildContext ...
func (ch *HttpChecker) BuildContext(cx model.CheckContext) {
	ch.Url = cx.Host
	ch.Status = cx.HttpConf.StatusCode
	if len(cx.HttpConf.Text) > 0 {
		ch.Text = cx.HttpConf.Text
	}
}

//Validate ...
func (ch *HttpChecker) Validate() error {
	if !ch.validateUrl() {
		return errors.New("invalid URL provided")
	}

	if !ch.validateStatusCode() {
		return errors.New("invalid status code provided")
	}

	return nil
}

//Check ...
func (ch *HttpChecker) Check() (bool, bool, error) {
	resp, err := http.Get(ch.Url)

	if err != nil {
		return false, true, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, true, err
	}

	if resp.StatusCode != ch.Status {
		return false, false, errors.New("invalid status code")
	}

	if len(ch.Text) > 0 {
		if !strings.Contains(string(body), ch.Text) {
			return false, false, errors.New("can't find substring in response")
		}
	}

	return true, false, nil
}

func (ch *HttpChecker) validateUrl() bool {
	_, err := url.ParseRequestURI(ch.Url)
	if err != nil {
		return false
	}

	u, err := url.Parse(ch.Url)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func (ch *HttpChecker) validateStatusCode() bool {
	// check against common status code
	if ch.Status < 100 || ch.Status > 599 {
		return false
	}
	return true
}
