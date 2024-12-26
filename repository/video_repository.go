package repository

import (
	"TejasThombare20/fampay/model"
	"context"
	"database/sql"
	"log"
	"time"
)

type VideoRepository struct {
	db *sql.DB
}

func NewVideoRepository(db *sql.DB) *VideoRepository {
	return &VideoRepository{db: db}
}

func (r *VideoRepository) Create(video *model.Video) error {
	query := `
        INSERT INTO videos (id, title, description, published_at, thumbnail_url, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (id) DO NOTHING`

	_, err := r.db.Exec(query,
		video.ID,
		video.Title,
		video.Description,
		video.PublishedAt,
		video.ThumbnailURL,
		time.Now(),
	)

	return err
}

func (r *VideoRepository) List(ctx context.Context, page, pageSize int) ([]model.Video, error) {
	offset := (page - 1) * pageSize

	query := `
        SELECT id, title, description, published_at, thumbnail_url, created_at
        FROM videos
        ORDER BY published_at DESC
        LIMIT $1 OFFSET $2`

	rows, err := r.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		log.Println("error while fetching the yt videos data", err)
		return nil, err
	}
	defer rows.Close()

	var videos []model.Video
	for rows.Next() {
		var video model.Video
		err := rows.Scan(
			&video.ID,
			&video.Title,
			&video.Description,
			&video.PublishedAt,
			&video.ThumbnailURL,
			&video.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return videos, nil
}
