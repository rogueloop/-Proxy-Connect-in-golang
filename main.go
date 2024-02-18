package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func ProxyAwareHttpClient() (*http.Client, error) {
	proxyURL, err := url.Parse("http://ProxyServerName.com/")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid proxy url:", err)
		return nil, err
	}

	proxyURL.User = url.UserPassword("usr_name_feild", "password_feild")

	// Setup an HTTP client with a proxy transport
	proxyTransport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	httpClient := &http.Client{Transport: proxyTransport}
	return httpClient, nil
}

func main() {
	client, err := ProxyAwareHttpClient()
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(contents))
}
