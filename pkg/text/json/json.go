package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type DocsList struct {
	Data []Doc `json:"data"`
}
type LastEditor struct {
	ID             int       `json:"id"`
	Type           string    `json:"type"`
	Login          string    `json:"login"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	AvatarURL      string    `json:"avatar_url"`
	FollowersCount int       `json:"followers_count"`
	FollowingCount int       `json:"following_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Serializer     string    `json:"_serializer"`
}
type Doc struct {
	ID                int         `json:"id"`
	Slug              string      `json:"slug"`
	Title             string      `json:"title"`
	Description       string      `json:"description"`
	UserID            int         `json:"user_id"`
	BookID            int         `json:"book_id"`
	Format            string      `json:"format"`
	Public            int         `json:"public"`
	Status            int         `json:"status"`
	ViewStatus        int         `json:"view_status"`
	ReadStatus        int         `json:"read_status"`
	LikesCount        int         `json:"likes_count"`
	CommentsCount     int         `json:"comments_count"`
	ContentUpdatedAt  time.Time   `json:"content_updated_at"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	PublishedAt       time.Time   `json:"published_at"`
	FirstPublishedAt  time.Time   `json:"first_published_at"`
	DraftVersion      int         `json:"draft_version"`
	LastEditorID      int         `json:"last_editor_id"`
	WordCount         int         `json:"word_count"`
	Cover             string      `json:"cover"`
	CustomDescription interface{} `json:"custom_description"`
	LastEditor        LastEditor  `json:"last_editor"`
	Book              interface{} `json:"book"`
	Serializer        string      `json:"_serializer"`
}

func main() {
	var a DocsList

	fileByte, err := os.ReadFile("./text_handle/test_files/test.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileByte, &a)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(a.Data))
}
