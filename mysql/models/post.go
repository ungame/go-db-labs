package models

import "time"

type Post struct {
	ID        int64
	AuthorID  string
	Title     string
	Body      string
	Likes     int64
	Dislikes  int64
	Favorites int64
	Deleted   byte
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
