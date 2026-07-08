package services

import (
	"blogger-api/models"
	"blogger-api/repositories"
)

type BlogService struct {
	repo *repositories.BlogRepository
}

func NewBlogService(repo *repositories.BlogRepository) *BlogService {
	return &BlogService{repo: repo}
}

func (s *BlogService) CreateBlog(title, author, body string) (models.Blog, error) {
	blog := models.Blog{Title: title, Author: author, Body: body}
	err := s.repo.Create(&blog)
	return blog, err
}

func (s *BlogService) GetAllBlogs() ([]models.Blog, error) {
	return s.repo.GetAll()
}

func (s *BlogService) GetBlogById(id uint) (models.Blog, error) {
	return s.repo.GetById(id)
}

func (s *BlogService) UpdateBlog(id uint, title, author, body string) (models.Blog, error) {
	blog, err := s.repo.GetById(id)
	if err != nil {
		return blog, err
	}

	blog.Title = title
	blog.Author = author
	blog.Body = body
	
	err = s.repo.Update(&blog)
	return blog, err
}

func (s *BlogService) DeleteBlog(id uint) error {
	return s.repo.Delete(id)
}