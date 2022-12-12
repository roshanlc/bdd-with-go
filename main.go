package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// constants
const randomQuoteURL = "https://stoicquotesapi.com/v1/api/quotes/random"

// A struct to hold all the quote details
// offered by the stoic api
type Quote struct {
	ID       int    `json:"id"`
	Body     string `json:"body"`
	AuthorID int    `json:"author_id"`
	Author   string `json:"author"`
}

// Bind String method to Quote
func (q *Quote) String() string {
	return fmt.Sprintf("Id=%d, author_id=%d, author=%#v , body=%#v\n", q.ID, q.AuthorID, q.Author, q.Body)
}

func main() {

	// Fetch a random quote
	quote, err := getRandomQuote()
	// error checking
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(quote)
}

// getRandomQuote returns Quote
func getRandomQuote() (*Quote, error) {
	resp := getReq(randomQuoteURL)
	if resp.err != nil {
		return nil, resp.err
	}
	// empty object
	var quote *Quote

	// defer close reading body of response
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal into quote
	err = json.Unmarshal(data, &quote)
	if err != nil {
		return nil, err
	}

	return quote, nil
}
