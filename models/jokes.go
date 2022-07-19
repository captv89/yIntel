package models

type Joke struct {
	Error    bool  `json:"error"`
	Category string `json:"category"`
	JokeType string `json:"type"`
	Joke     []string `json:"joke"`
	Flags    map[string]bool `json:"flags"`
	Id       float64 `json:"id"`
	Lang     string `json:"lang"`
}

type Chuck struct {
	Category []string `json:"category"`
	CreatedAt string `json:"created_at"`
	IconUrl  string `json:"icon_url"`
	Id       string `json:"id"`
	UpdatedAt string `json:"updated_at"`
	Url      string `json:"url"`
	Value    string `json:"value"`
}

type DadJoke struct {
	ID        string `json:"id"`
	Joke      string `json:"joke"`
	Status    int `json:"status"`
}

type Meme struct {
	PostLink string `json:"postLink"`
	SubReddit string `json:"subreddit"`
	Title    string `json:"title"`
	Upvotes  int `json:"ups"`
	Author   string `json:"author"`
	URL   string `json:"url"`
}