package core

//http://www.01happy.com/golang-http-client-get-and-post/

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body)
}

func HttpPost(url string) string {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(""))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body)
}

func httpPostForm(url string) string {
	resp, err := http.PostForm(url,
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body)
}

func httpRequest(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	return string(body)
}

func HttpPut() {

}

func HttpDelete() {

}

func HttpHead() {

}
