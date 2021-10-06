package entry

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Page struct {
	Name    string
	Entries []Entry
}

type Entry struct {
	ID    string
	Title string
	Body  string
	Date  string
	Tags  []string
}

func NewEntry(ID string, Title string, Body string, Date string, Tags []string) *Entry {
	return &Entry{
		ID:    ID,
		Title: Title,
		Body:  Body,
		Date:  Date,
		Tags:  Tags,
	}
}

func (p *Entry) SaveEntry() error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	filename := p.ID + ".json"
	b, _ := json.Marshal(p)

	return os.WriteFile("data/"+filename, b, 0600)
}

func LoadSingleEntry(id string) (*Entry, error) {
	filename := id + ".json"
	guts, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var post Entry
	err = json.Unmarshal(guts, &post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func LoadEntry(fileName string) (*Entry, error) {
	data, err := os.ReadFile("data/" + fileName)
	if err != nil {
		return nil, err
	}
	var entry Entry
	err = json.Unmarshal(data, &entry)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// fmt.Printf("the entry %s", entry)
	// fmt.Println("**********************")

	return &entry, nil
}

func LoadEntries() (*[]Entry, error) {
	// println("In LoadEntries!")
	files, err := os.ReadDir("data")
	println(len(files))
	if err != nil {
		return nil, err
	}
	entries := []Entry{}

	for _, file := range files {
		if !file.IsDir() {
			entry, _ := LoadEntry(file.Name())
			entries = append(entries, *entry)
		}
	}
	return &entries, nil
}
