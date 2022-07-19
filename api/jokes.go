package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/captv89/yIntel/models"
)

func Jokes() (models.Joke, bool) {
	url := "https://v2.jokeapi.dev/joke/Any?type=single"
	resp, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var joke models.Joke

	// Check the status code
	if resp.StatusCode != 200 {
		log.Println("No Jokes, Status code is not 200")
		return joke, false
	}
	// Print response body
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(joke.Joke)
	return joke, true
}

// Get chucknorris joke
func Chuck() (models.Chuck, bool) {
	url := "https://api.chucknorris.io/jokes/random"
	resp, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var joke models.Chuck

	// Check the status code
	if resp.StatusCode != 200 {
		log.Println("No Jokes, Status code is not 200")
		return joke, false
	}
	// Print response body
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Category: ",strings.Join(joke.Category, ","))

	// fmt.Println(joke.Value)
	return joke, true
}

func DadJokes() (models.DadJoke, bool) {
	url := "https://icanhazdadjoke.com/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := myClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var joke models.DadJoke

	// Check the status code
	if resp.StatusCode != 200 {
		log.Println("No Jokes, Status code is not 200")
		return joke, false
	}

	// Write joke to the body
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		log.Fatal(err)
	}

	return joke, true
}


//  Get meme

func Meme() (models.Meme, bool) {
	url := "https://meme-api.herokuapp.com/gimme"

	resp, err := myClient.Get(url)
	if err != nil {	
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var meme models.Meme

	// Check the status code
	if resp.StatusCode != 200 {
		log.Println("No Jokes, Status code is not 200")
		return meme, false
	}

	// Write joke to the body
	err = json.NewDecoder(resp.Body).Decode(&meme)
	if err != nil {
		log.Fatal(err)
	}

	return meme, true
}
