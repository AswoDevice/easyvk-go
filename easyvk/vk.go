// Package easyvk provides you simple way
// to work with VK API.
package easyvk

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

const (
	version = "5.92"
	apiURL  = "https://api.vk.com/method/"
	authURL = "https://oauth.vk.com/authorize?" +
		"client_id=%s" +
		"&scope=%s" +
		"&redirect_uri=https://oauth.vk.com/blank.html" +
		"&display=wap" +
		"&v=%s" +
		"&response_type=token"
	authClientFlowURL = "https://oauth.vk.com/access_token?" +
		"client_id=%s" +
		"&client_secret=%s" +
		"&v=%s" +
		"&grant_type=client_credentials"
)

// VK defines a set of functions for
// working with VK API.
type VK struct {
	AccessToken   string
	Version       string
	Account       Account
	Board         Board
	Database      Database
	Fave          Fave
	Friends       Friends
	Groups        Groups
	Likes         Likes
	Messages      Messages
	Notifications Notifications
	Photos        Photos
	Status        Status
	Storage       Storage
	Upload        Upload
	Users         Users
	Wall          Wall
}

// WithToken helps to initialize your
// VK object with token.
func WithToken(token string) VK {
	vk := VK{}
	vk.AccessToken = token
	vk.Version = version
	vk.Account = Account{&vk}
	vk.Board = Board{&vk}
	vk.Database = Database{&vk}
	vk.Fave = Fave{&vk}
	vk.Friends = Friends{&vk}
	vk.Groups = Groups{&vk}
	vk.Likes = Likes{&vk}
	vk.Messages = Messages{&vk}
	vk.Notifications = Notifications{&vk}
	vk.Photos = Photos{&vk}
	vk.Status = Status{&vk}
	vk.Storage = Storage{&vk}
	vk.Upload = Upload{}
	vk.Users = Users{&vk}
	vk.Wall = Wall{&vk}
	return vk
}

// WithAuth helps to initialize your VK object
// with signing in by login, password, client id and scope
// Scope must be a string like "friends,wall"
func WithAuth(login, password, clientID, scope string) (VK, error) {
	u := fmt.Sprintf(authURL, clientID, scope, version)
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	resp, err := client.Get(u)
	if err != nil {
		return VK{}, err
	}
	defer resp.Body.Close()

	args, u := parseForm(resp.Body)

	args.Add("email", login)
	args.Add("pass", password)

	resp, err = client.PostForm(u, args)
	if err != nil {
		return VK{}, err
	}

	if resp.Request.URL.Path != "/blank.html" {
		args, u := parseForm(resp.Body)
		resp, err := client.PostForm(u, args)
		if err != nil {
			return VK{}, err
		}
		defer resp.Body.Close()

		if resp.Request.URL.Path != "/blank.html" {
			return VK{}, errors.New("can't log in")
		}
	}

	urlArgs, err := url.ParseQuery(resp.Request.URL.Fragment)
	if err != nil {
		return VK{}, err
	}

	return WithToken(urlArgs["access_token"][0]), nil
}

func WithClientFlow(clientID, clientSecret string) (VK, error) {
	u := fmt.Sprintf(authClientFlowURL, clientID, clientSecret, version)

	resp, err := http.DefaultClient.Get(u)
	if err != nil {
		return VK{}, err
	}
	defer resp.Body.Close()

	type Response struct {
		AccessToken string `json:"access_token"`
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return VK{}, err
	}

	return WithToken(response.AccessToken), nil
}

func parseForm(body io.ReadCloser) (url.Values, string) {
	tokenizer := html.NewTokenizer(body)

	u := ""
	formData := map[string]string{}

	end := false
	for !end {
		tag := tokenizer.Next()

		switch tag {
		case html.ErrorToken:
			end = true
		case html.StartTagToken:
			switch token := tokenizer.Token(); token.Data {
			case "form":
				for _, attr := range token.Attr {
					if attr.Key == "action" {
						u = attr.Val
					}
				}
			case "input":
				if token.Attr[1].Val == "_origin" {
					formData["_origin"] = token.Attr[2].Val
				}
				if token.Attr[1].Val == "to" {
					formData["to"] = token.Attr[2].Val
				}
			}
		case html.SelfClosingTagToken:
			switch token := tokenizer.Token(); token.Data {
			case "input":
				if token.Attr[1].Val == "ip_h" {
					formData["ip_h"] = token.Attr[2].Val
				}
				if token.Attr[1].Val == "lg_h" {
					formData["lg_h"] = token.Attr[2].Val
				}
			}
		}
	}

	args := url.Values{}

	for key, val := range formData {
		args.Add(key, val)
	}

	return args, u
}

// Request provides access to VK API methods.
func (vk *VK) Request(method string, params map[string]string) (data []byte, err error) {
	for attempt := 0; attempt < 10; attempt++ {
		data, err = vk.RequestAttempt(method, params)
		if err == nil {
			return data, err
		}
		duration :=
			(500 * time.Millisecond) +
				(time.Millisecond * time.Duration(rand.Int31n(2000)))
		time.Sleep(duration)
	}

	return data, err
}

// Request provides access to VK API methods.
func (vk *VK) RequestAttempt(method string, params map[string]string) ([]byte, error) {
	u, err := url.Parse(apiURL + method)
	if err != nil {
		return nil, err
	}

	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	query.Set("access_token", vk.AccessToken)
	query.Set("v", vk.Version)
	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var handler struct {
		Error    *Error
		Response json.RawMessage
	}
	err = json.Unmarshal(body, &handler)

	if handler.Error != nil {
		return nil, handler.Error
	}

	return handler.Response, nil
}
