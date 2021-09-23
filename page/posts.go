package page

import (
	"encoding/json"
	"os"
	"time"
)

type Page struct {
	Posts []Post
}

type Post struct {
	ID    string
	Title string
	Body  []byte
	Date  time.Time
	Tags  []string
}

func (p *Post) savePost() error {
	filename := p.ID + ".json"
	b, _ := json.Marshal(p)

	return os.WriteFile(filename, b, 0600)
}

func loadPost(id string) (*Post, error) {
	filename := id + ".json"
	guts, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var post Post
	err = json.Unmarshal(guts, &post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
