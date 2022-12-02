package post

import (
	context "context"
	"database/sql"
	"iman/pkg/proto/post_service"
)

type Repository interface {
	GetPosts(ctx context.Context, limit, offset int) (posts []*post_service.Post, err error)
	GetPostByID(ctx context.Context, id int) (*post_service.Post, error)
	DeletePostByID(ctx context.Context, id int) (err error)
	UpdatePostByID(ctx context.Context, id int, title, body string) (err error)
	GetPostsCount(ctx context.Context) (cnt int, err error)
	CreatePost(ctx context.Context, post *post_service.Post) (err error)
}

type repository struct {
	DB *sql.DB
}

type Params struct {
	DB *sql.DB
}

func New(p Params) Repository {
	return &repository{
		DB: p.DB,
	}
}

func (r *repository) GetPosts(ctx context.Context, limit, offset int) (posts []*post_service.Post, err error) {
	query := `
		SELECT id, user_id, title, body
		FROM posts
		ORDER BY id DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p post_service.Post
		err = rows.Scan(&p.Id, &p.UserId, &p.Title, &p.Body)
		if err != nil {
			return
		}

		posts = append(posts, &p)
	}

	return
}

func (r *repository) GetPostByID(ctx context.Context, id int) (*post_service.Post, error) {
	query := `
		SELECT id, user_id, title, body
		FROM posts
		WHERE id = $1
	`

	var post post_service.Post

	err := r.DB.QueryRow(query, id).Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *repository) DeletePostByID(ctx context.Context, id int) (err error) {
	query := `
		DELETE FROM posts
		WHERE id = $1
	`

	res, err := r.DB.Exec(query, id)
	if err != nil {
		return
	}

	cnt, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if cnt == 0 {
		return sql.ErrNoRows
	}

	return
}

func (r *repository) UpdatePostByID(ctx context.Context, id int, title, body string) (err error) {
	query := `
		UPDATE posts
		SET title = $1, body = $2
		WHERE id = $3
	`

	res, err := r.DB.Exec(query, title, body, id)
	if err != nil {
		return err
	}

	cnt, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if cnt == 0 {
		return sql.ErrNoRows
	}

	return
}

func (r *repository) GetPostsCount(ctx context.Context) (cnt int, err error) {
	query := `
		SELECT count(*)
		FROM posts
	`

	err = r.DB.QueryRow(query).Scan(&cnt)
	if err != nil {
		return
	}

	return
}

func (r *repository) CreatePost(ctx context.Context, post *post_service.Post) (err error) {
	query := `
		INSERT INTO posts(id, user_id, title, body)
		VALUES($1, $2, $3, $4)
	`

	_, err = r.DB.Exec(query, post.Id, post.UserId, post.Title, post.Body)
	if err != nil {
		return
	}

	return
}
