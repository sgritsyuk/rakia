package store

import (
	"api-service/internal/domain"
)

// PostEntry represent database document structure
type PostEntry struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// convert entry to domain structure
func (p *PostEntry) toDomain() domain.Post {
	return domain.Post{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
		Author:  p.Author,
	}
}
