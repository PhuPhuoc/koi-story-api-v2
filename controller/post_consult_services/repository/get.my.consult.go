package consultrepository

import (
	"fmt"

	consultmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_consult_services/model"
	"github.com/jmoiron/sqlx"
)

func (store *consultStore) GetMyListPostConsult(user_id string) ([]consultmodel.PostConsult, error) {
	list_consult_post := []consultmodel.PostConsult{}
	tx, err := store.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("cannot begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p) // Rethrow panic after rollback.
		} else if err != nil {
			_ = tx.Rollback() // Ignore rollback error and return original error.
		} else if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("cannot commit transaction: %w", commitErr)
		}
	}()

	consult := []consultmodel.ConsultModelForGet{}

	query := `
		select c.post_id, c.title, c.content
		from
			consults c
		join
			posts p on p.id=c.post_id
		where p.user_id=?
	`
	if err := tx.Select(&consult, query, user_id); err != nil {
		return nil, fmt.Errorf("unable to get information of consult post from db: %w", err)
	}

	postIDs := []string{}
	for _, consult := range consult {
		postIDs = append(postIDs, consult.PostID)
	}

	// query get all image for all post
	images := []consultmodel.DetailImage{}
	query_image := `
        select id, image_url, post_id
        from images
        where post_id in (?)
    `
	query, args, err := sqlx.In(query_image, postIDs)
	if err != nil {
		return nil, fmt.Errorf("error preparing query: %w", err)
	}
	query = tx.Rebind(query)
	if err := tx.Select(&images, query, args...); err != nil {
		return nil, fmt.Errorf("unable to get images of consult's post from db: %w", err)
	}
	imgMap := make(map[string][]consultmodel.DetailImage)
	for _, image := range images {
		imgMap[image.PostID] = append(imgMap[image.PostID], image)
	}

	comments := []consultmodel.CommentInPostConsult{}
	query_comment := `
        select
            c.id, c.post_id, c.user_id, c.content, c.created_at, u.name, u.avatar
        from
            comments c
        join
            users u on c.user_id=u.id
        where
            c.post_id in (?)
        order by
            c.created_at desc
        limit 2
    `
	query_cmt, args_cmt, err_cmt := sqlx.In(query_comment, postIDs)
	if err_cmt != nil {
		return nil, fmt.Errorf("error preparing query: %w", err)
	}
	query_cmt = tx.Rebind(query_cmt)
	if err := tx.Select(&comments, query_cmt, args_cmt...); err != nil {
		return nil, fmt.Errorf("unable to get comments of consult's post from db: %w", err)
	}

	cmtMap := make(map[string][]consultmodel.CommentInPostConsult)
	for _, cmt := range comments {
		cmtMap[cmt.PostID] = append(cmtMap[cmt.PostID], cmt)
	}

	for _, post := range consult {
		fullpost := consultmodel.PostConsult{
			ConsultModelForGet: post,
			Images:             imgMap[post.PostID],
			Comments:           cmtMap[post.PostID],
		}
		list_consult_post = append(list_consult_post, fullpost)
	}

	return list_consult_post, nil
}
