package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

// Story is a struct that represents a Hacker News story
type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

var stories []Story

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the top news from the API
	resp, err := myClient.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Print response body
	var topStories []int
	err = json.NewDecoder(resp.Body).Decode(&topStories)

	fmt.Println(topStories)

	if err != nil {
		log.Fatal(err)
	}

	// Print the top stories
	for _, story := range topStories {
		storyUrl := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", story)
		log.Println("Url:",storyUrl)
		resp, err = myClient.Get(storyUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		//  Print response body
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Println("Body:", string(body))

		// Get the response code
		log.Println("Status:", resp.StatusCode)

		// unmarshal the response into a Story struct
		var storyData Story
		err = json.NewDecoder(resp.Body).Decode(&storyData)
		if err != nil {
			log.Fatal(err)
		}

		// Print the title of the story
		fmt.Println(storyData.Title)

		// Append the story to the stories slice
		stories = append(stories, storyData)
	}

	// Write the stories to a JSON file
	jsonData, err := json.Marshal(stories)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("stories.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

}