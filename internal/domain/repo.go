package domain

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/yumubi/bookmarks.git/internal/config"
	"time"
)

type Bookmark struct {
	ID        int
	Title     string
	Url       string
	CreatedAt time.Time
}

type BookmarkRepository interface {
	GetAll(ctx context.Context) ([]Bookmark, error)
	GetByID(ctx context.Context, id int) (*Bookmark, error)
	Create(ctx context.Context, b Bookmark) (*Bookmark, error)
	Update(ctx context.Context, b Bookmark) error
	Delete(ctx context.Context, id int) error
}

type bookmarkRepo struct {
	db     *pgx.Conn
	logger *config.Logger
}

func (r bookmarkRepo) GetAll(ctx context.Context) ([]Bookmark, error) {
	query := `select id, title, url, created_at FROM bookmarks`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var bookmarks []Bookmark
	for rows.Next() {
		var bookmark = Bookmark{}
		err := rows.Scan(&bookmark.ID, &bookmark.Title, &bookmark.Url, &bookmark.CreatedAt)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return bookmarks, nil

}

func (b bookmarkRepo) GetByID(ctx context.Context, id int) (*Bookmark, error) {
	//TODO implement me
	panic("implement me")
}

func (r bookmarkRepo) Create(ctx context.Context, b Bookmark) (*Bookmark, error) {
	//TODO implement me
	panic("implement me")
}

func (r bookmarkRepo) Update(ctx context.Context, b Bookmark) error {
	//TODO implement me
	panic("implement me")
}

func (r bookmarkRepo) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func NewBookRepository(db *pgx.Conn, logger *config.Logger) BookmarkRepository {
	return bookmarkRepo{
		db:     db,
		logger: logger,
	}

}
