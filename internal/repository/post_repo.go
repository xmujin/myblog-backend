package repository

import (
	"errors"

	"github.com/xmujin/myblog-backend/internal/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(post *model.Post) error
	GetPosts() ([]model.Post, error)
	GetPostById(id uint) (*model.Post, error)
	DeletePostById(id uint) error
	UpdatePost(post *model.Post) error
}

type postgresPostRepository struct {
	db *gorm.DB
}

func (p *postgresPostRepository) UpdatePost(post *model.Post) error {
	return p.db.Save(post).Error
}

func (p *postgresPostRepository) DeletePostById(id uint) error {
	rowsAffected := p.db.Delete(&model.Post{}, id).RowsAffected
	if rowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}

func (p *postgresPostRepository) GetPostById(id uint) (*model.Post, error) {
	var post model.Post
	err := p.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postgresPostRepository{
		db: db,
	}
}

func (p *postgresPostRepository) CreatePost(post *model.Post) error {
	return p.db.Create(post).Error
}

func (p *postgresPostRepository) GetPosts() ([]model.Post, error) {
	var posts []model.Post
	err := p.db.Find(&posts).Error

	if err != nil {
		return nil, err
	}
	return posts, nil
}
