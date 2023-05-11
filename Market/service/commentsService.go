package service

import (
	"github.com/Moldaspan/E-commerce/models"
	"github.com/Moldaspan/E-commerce/repositories"
)

type CommentServiceInterface interface {
	CreateComment(*models.Comment) error
	GetCommentById(uint) (*models.Comment, error)
	DeleteComment(uint) error
}

type CommentServiceV1 struct {
	commentRepos repositories.CommentRepositoryInterface
}

func NewCommentService() CommentServiceV1 {
	return CommentServiceV1{commentRepos: repositories.NewCommentRepository()}
}

func (c CommentServiceV1) CreateComment(comment *Comment) error {
	return c.commentRepos.CreateComment(comment)
}

func (c CommentServiceV1) GetCommentById(id uint) (*Comment, error) {
	return c.commentRepos.GetCommentById(id)
}
