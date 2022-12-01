package post

import (
	context "context"
	"database/sql"
	"iman/pkg/proto/post_service"
)

type Repository struct {
	DB *sql.DB
}

func (p *Repository) GetPosts(ctx context.Context, limit, offset int) (posts []*post_service.Post, err error) {
	query := `
		SELECT id, user_id, title, body
		FROM posts
		LIMIT $1 OFFSET $2
	`

	rows, err := p.DB.Query(query, limit, offset)
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

func (p *Repository) GetPostByID(ctx context.Context, id int) (*post_service.Post, error) {
	query := `
		SELECT id, user_id, title, body
		FROM posts
		WHERE id = $1
	`

	var post post_service.Post

	err := p.DB.QueryRow(query, id).Scan(&post.Id, &post.UserId, &post.Title, &post.Body)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *Repository) DeletePostByID(ctx context.Context, id int) (err error) {
	query := `
		DELETE FROM posts
		WHERE id = $1
	`

	res, err := p.DB.Exec(query, id)
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

func (p *Repository) UpdatePostByID(ctx context.Context, id int, title, body string) (err error) {
	query := `
		UPDATE posts
		SET title = $1, body = $2
		WHERE id = $3
	`

	res, err := p.DB.Exec(query, title, body, id)
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

func (p *Repository) GetPostsCount(ctx context.Context) (cnt int, err error) {
	query := `
		SELECT count(*)
		FROM posts
	`

	err = p.DB.QueryRow(query).Scan(&cnt)
	if err != nil {
		return
	}

	return
}

func (p *Repository) CreatePost(ctx context.Context, post *post_service.Post) (err error) {
	query := `
		INSERT INTO posts(id, user_id, title, body)
		VALUES($1, $2, $3, $4)
	`

	_, err = p.DB.Exec(query, post.Id, post.UserId, post.Title, post.Body)
	if err != nil {
		return
	}

	return
}
