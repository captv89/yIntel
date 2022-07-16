package ops

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/captv89/yIntel/models"
)

// Read top stories from the json file
func ReadStory() []models.Story {
	stories := []models.Story{}

	// Read the json file
	file, err := ioutil.ReadFile("stories.json")
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the json file into the stories slice
	err = json.Unmarshal([]byte(file), &stories)
	if err != nil {
		log.Fatal(err)
	}

	// // Print the top stories
	// for _, story := range stories {
	// 	fmt.Println(story.Title)
	// }

	return stories
}

func Top10Stories() []string {
	// Read the top stories from the json file
	stories := ReadStory()

	tenNews := []string{}

	// Get the tenNews of the top stories
	for i := 0; i < 10; i++ {
		newsTime := time.Unix(int64(stories[i].Time), 0)
		t := newsTime.Format("2006-01-02 15:04:05")
		newsItem := fmt.Sprintf("%d. Title: %s \n Score: %d \n Time: %s \n Type: %s \n Link: %s \n", i+1, stories[i].Title, stories[i].Score, t, stories[i].Type, stories[i].URL)

		tenNews = append(tenNews, newsItem)
	}

	return tenNews
	// // Convert the tenNews slice to a string
	// tenNewsString := fmt.Sprintf("%s", tenNews)

	// return tenNewsString
}

// RandomStory is a function that returns a random story from the top stories
func RandomStory() string {
	// Read the top stories from the json file
	stories := ReadStory()

	// Get a random story
	rand.Seed(time.Now().UnixNano())
	randomStory := stories[rand.Intn(len(stories))]

	// Time
	newsTime := time.Unix(int64(randomStory.Time), 0)
	t := newsTime.Format("2006-01-02 15:04:05")

	// Convert the random story to a string
	randomStoryString := fmt.Sprintf("Title: %s \n Score: %d \n Time: %s \n Type: %s \n Link: %s \n", randomStory.Title, randomStory.Score, t, randomStory.Type, randomStory.URL)

	return randomStoryString
}
