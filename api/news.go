package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/captv89/yIntel/models"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

var stories []models.Story

func FetchStory() {
	// Get the top news from the API
	resp, err := myClient.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Print response body
	var topStories []int
	err = json.NewDecoder(resp.Body).Decode(&topStories)

	// fmt.Println(topStories)

	if err != nil {
		log.Fatal(err)
	}

	// Print the top stories
	for _, story := range topStories {
		storyUrl := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", story)
		log.Println("Url:", storyUrl)
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
		var storyData models.Story
		err = json.NewDecoder(resp.Body).Decode(&storyData)
		if err != nil {
			log.Fatal(err)
		}

		// Print the title of the story
		log.Println("Title:", storyData.Title)

		// Append the story to the stories slice
		stories = append(stories, storyData)
	}

	// Write the stories to a JSON file
	jsonData, err := json.Marshal(stories)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("stories.json", jsonData, 0o644)
	if err != nil {
		log.Fatal(err)
	}
}

// Fetch the top news from the API every day at midnight
func init() {
	go func() {
		for {
			// check stories.json last modified time
			log.Println("Checking stories.json last modified time")
			currentTime := time.Now()
			file, err := os.Stat("stories.json")
			if err != nil {
				log.Fatal(err)
			}
			lastModifiedTime := file.ModTime()
			// if stories.json is older than 24 hours, fetch the top news
			if currentTime.Sub(lastModifiedTime) > 24*time.Hour {
				FetchStory()
			} else {
				log.Println("stories.json is less than 24 hours old")
			}
			// wait for 24 hours
			log.Println("Waiting for 24 hours")
			time.Sleep(24 * time.Hour)
		}
	}()
}


