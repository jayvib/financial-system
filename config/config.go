package config

import (
	"golang.org/x/oauth2"
	"fmt"
	"log"
	"os"
	"encoding/json"
	"context"
	"net/http"
	"io/ioutil"
	"golang.org/x/oauth2/google"
)

type configFunc func([]byte) (*oauth2.Config, error)


// Retrieve a token, saves the token, then returns the generated client.
func NewClient(filename string, confFunc configFunc) (*http.Client, error) {
	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Printf("Unable to read client secret file: %v", err)
		return nil, err
	}

	conf, err := confFunc(b)
	if err != nil {
		return nil, err
	}

	client := getClient(conf)
	return client, nil
}


// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func ReadOnlyConfigFunc(b []byte) (*oauth2.Config, error) {
	conf, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func ReadWriteConfigFunc(b []byte) (*oauth2.Config, error) {
	conf, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}
	return conf, nil
}