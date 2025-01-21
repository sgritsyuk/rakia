package store

import (
	"api-service/internal/domain"
	"context"
	"encoding/json"
	"io"
	"os"
	"sort"
	"strings"
)

type FileData struct {
	Posts []PostEntry `json:"posts"`
}

// MemoryPostStore allows to store and retrieve posts
type MemoryPostStore struct {
	collection    map[int]PostEntry
	autoincrement int
}

// NewMemoryPostStore creates a new implementation of posts store
func NewMemoryPostStore(initFile string) (*MemoryPostStore, error) {
	file, err := os.Open(initFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var data FileData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	var collection = make(map[int]PostEntry)
	var maxID int = 0
	for _, post := range data.Posts {
		if post.ID > maxID {
			maxID = post.ID
		}
		collection[post.ID] = post
	}
	return &MemoryPostStore{
		collection:    collection,
		autoincrement: maxID,
	}, nil
}

// Get fetch the list of posts according specified criteria (inc pagination)
func (s *MemoryPostStore) Get(ctx context.Context, title string, page int, limit int) (*[]domain.Post, error) {
	// filter data
	titleLower := strings.ToLower(title)
	arr := make([]domain.Post, 0, len(s.collection))
	for _, value := range s.collection {
		if len(title) == 0 || strings.Contains(strings.ToLower(value.Title), titleLower) {
			arr = append(arr, value.toDomain())
		}
	}
	// sort by id
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].ID < arr[j].ID
	})
	// apply page, limit
	start := (page - 1) * limit
	end := start + limit
	if start > len(arr) {
		return &[]domain.Post{}, nil
	}
	if end > len(arr) {
		end = len(arr)
	}
	paginated := arr[start:end]
	return &paginated, nil
}

// GetOne fetch the one post according to specified id
func (s *MemoryPostStore) GetOne(ctx context.Context, id int) (*domain.Post, error) {
	doc, ok := s.collection[id]
	if !ok {
		return &domain.Post{}, domain.ErrorPostNotFound
	}
	post := doc.toDomain()
	return &post, nil
}

func (s *MemoryPostStore) Insert(ctx context.Context, post domain.Post) (int, error) {
	// increase ID counter
	s.autoincrement++
	// construct document structure
	doc := PostEntry{
		ID:      s.autoincrement,
		Title:   post.Title,
		Content: post.Content,
		Author:  post.Author,
	}
	// insert document into storage
	s.collection[doc.ID] = doc
	return doc.ID, nil
}

func (s *MemoryPostStore) Update(ctx context.Context, id int, post domain.Post) error {
	// check id exists
	doc, ok := s.collection[id]
	if !ok {
		return domain.ErrorPostNotFound
	}
	// update document
	doc.Title = post.Title
	doc.Content = post.Content
	doc.Author = post.Author
	s.collection[id] = doc
	return nil
}

func (s *MemoryPostStore) Delete(ctx context.Context, id int) error {
	// check id exists
	_, ok := s.collection[id]
	if !ok {
		return domain.ErrorPostNotFound
	}
	// delete document
	delete(s.collection, id)
	return nil
}
