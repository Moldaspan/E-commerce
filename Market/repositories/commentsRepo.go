package repositories

import (
	"github.com/Moldaspan/E-commerce/settings"
	"gorm.io/gorm"
	"log"
)

type CommentRepositoryInterface interface {
	CreateComment(*Comment) error
	GetCommentById(uint) (*Comment, error)
	DeleteComment(uint) error
}

type CommentRepositoryV1 struct {
	DB *gorm.DB
}

func NewCommentRepository() *CommentRepositoryV1 {
	db, err := settings.DbSetup()
	if err != nil {
		log.Fatal(err)
		return &CommentRepositoryV1{}
	}
	return &CommentRepositoryV1{DB: db}
}

func (c *CommentRepositoryV1) CreateComment(comment *Comment) error {
	return c.DB.Create(comment).Error
}

func (c *CommentRepositoryV1) DeleteComment(id uint) error {
	return c.DB.Delete(&Comment{}, id).Error
}

func (c *CommentRepositoryV1) GetCommentById(id uint) (*Comment, error) {
	var comment Comment
	if err := c.DB.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}
