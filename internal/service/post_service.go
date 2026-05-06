package service

import (
	"github.com/xmujin/myblog-backend/internal/model"
	"github.com/xmujin/myblog-backend/internal/repository"
)

type PostService interface {
	CreatePost(request *model.PostRequest) error
	GetPosts() ([]model.Post, error)
	GetPost(id uint) (*model.Post, error)
	DeletePostById(id uint) error
	UpdatePostById(id uint, request *model.PostRequest) error
}

type service struct {
	postRepository repository.PostRepository
}

func (s *service) UpdatePostById(id uint, request *model.PostRequest) error {
	post, err := s.postRepository.GetPostById(id)
	if err != nil {
		return err
	}
	post.Title = request.Title
	post.Content = request.Content
	return s.postRepository.UpdatePost(post)
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &service{
		postRepository: postRepository,
	}
}

func (s *service) DeletePostById(id uint) error {
	return s.postRepository.DeletePostById(id)
}

func (s *service) GetPost(id uint) (*model.Post, error) {
	post, err := s.postRepository.GetPostById(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *service) GetPosts() ([]model.Post, error) {
	posts, err := s.postRepository.GetPosts()

	// TODO 将model映射为dto
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *service) CreatePost(request *model.PostRequest) error {
	post := &model.Post{
		Title:   request.Title,
		Content: request.Content,
	}
	return s.postRepository.CreatePost(post)
}
