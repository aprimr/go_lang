package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {

	fmt.Println(myUrl)
	res, _ := url.Parse(myUrl)

	// fmt.Println(res.Scheme)
	// fmt.Println(res.Host)
	// fmt.Println(res.Path)
	// fmt.Println(res.RawQuery)

	queryParams := res.Query()
	// fmt.Println(queryParams["v"])

	for _, val := range queryParams {
		fmt.Println(val)
	}

	// constructing url
	newUrl := &url.URL{
		Scheme:  res.Scheme,
		Host:    res.Host,
		Path:    res.Path,
		RawQuery: res.RawQuery,
	}

	fmt.Println("The url is ", newUrl.String())
}
