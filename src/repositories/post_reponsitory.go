package repositories

import (
	"giasuaeapi/src/entities"

	"gorm.io/gorm"
)

type PostRepository interface {
	InsertPost(p *entities.Post) error
	UpdatePost(p *entities.Post) error
	DeletePost(id int) error
	FindAllPost() []entities.Post
	FindByID(id int) entities.Post
	FilterPost(typeID int, page int, pageSize int) []entities.Post
}

type postConnection struct {
	connection *gorm.DB
}

// FilterPost implements PostRepository
func (db *postConnection) FilterPost(typeID int, page int, pageSize int) []entities.Post {
	var posts []entities.Post
	db.connection.Limit(pageSize).Offset((page-1)*pageSize).Where("type = ?", typeID).Find(&posts)
	return posts
}

// DeletePost implements PostRepository
func (db *postConnection) DeletePost(id int) error {
	err := db.connection.Where("id = ?", id).Delete(&entities.Post{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// FindAllPost implements PostRepository
func (db *postConnection) FindAllPost() []entities.Post {
	var postes []entities.Post
	db.connection.Find(&postes)
	return postes
}

// FindByID implements PostRepository
func (db *postConnection) FindByID(id int) entities.Post {
	var post entities.Post
	db.connection.First(&post, id)
	return post
}

// InsertPost implements PostRepository
func (db *postConnection) InsertPost(p *entities.Post) error {
	err := db.connection.Save(&p)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// UpdatePost implements PostRepository
func (db *postConnection) UpdatePost(p *entities.Post) error {
	err := db.connection.Save(&p)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func NewPostRepository(dbConn *gorm.DB) PostRepository {
	return &postConnection{
		connection: dbConn,
	}
}
