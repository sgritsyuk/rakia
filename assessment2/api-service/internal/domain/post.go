package domain

import (
	"errors"
)

// Post domain structure
type Post struct {
	ID      int
	Title   string
	Content string
	Author  string
}

// ErrorPostNotFound is returned by some functions when a post is not found
var ErrorPostNotFound = errors.New("post not found")
