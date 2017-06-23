package main

import (
	"fmt"
	"net/http"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/net/context"
	"bytes"
)

func handleMain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}

func main() {
	config := &clientcredentials.Config{
		ClientID:     "iXyhYOR6pJl9TtYyYSzoiBXiB6sa",
		ClientSecret: "6_9n4Hqz7j8jGQ0P6A6gjcGWXCka",
		Scopes:       []string{"default"},
		TokenURL:     "https://ide.slcq002.com/oauth2/token",
	}

	client := config.Client(context.Background())
	// the client will update its token if it's expired
	resp, err := client.Get("https://api-dev.slcq002.com/search/catalog/categories/v3/?id=174")
	if err != nil{
		panic(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()
	fmt.Println(newStr)
	fmt.Print("Started running on http://127.0.0.1:7000\n")
	//fmt.Println(http.ListenAndServe(":7000", nil))
}