package repo

import (
	"context"
	"crawler/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *RepoPG) InsertPost(ctx context.Context, post model.Post, tx *gorm.DB) error {
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}

	if err := tx.Model(&model.Post{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "url_origin"}},
		UpdateAll: true,
	}).Create(&post).Error; err != nil {
		return err
	}
	return nil
}
func (r *RepoPG) InsertMultiPosts(ctx context.Context, posts []model.Post, tx *gorm.DB) error {
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}
	if err := tx.Model(&model.Post{}).Create(&posts).Error; err != nil {
		return err
	}
	return nil
}
