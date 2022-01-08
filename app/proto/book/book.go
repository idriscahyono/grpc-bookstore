package book

import (
	"context"
	"github.com/idriscahyono/grpc-bookstore/app/configs"
	"github.com/idriscahyono/grpc-bookstore/app/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Server struct {
	UnimplementedBookServiceServer
}

func (s *Server) CreateBook(ctx context.Context, b *Book) (*Book, error) {
	log.Println("Create Book...")

	if &b == nil {
		return nil, status.Error(codes.InvalidArgument, "can't be empty")
	}

	if err := configs.DB.Create(b).Error; err != nil {
		return nil, status.Error(codes.InvalidArgument, "failed create book")
	}

	return b, nil
}

func (s *Server) ReadBook(ctx context.Context, b *ID) (*Book, error) {
	log.Println("Find Book...")

	if &b.ID == nil {
		return nil, status.Error(codes.InvalidArgument, "ID can't be empty")
	}

	var result Book
	if err := configs.DB.Where("ID = ?", b.ID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Server) UpdateBook(ctx context.Context, b *Book) (*Book, error) {
	var bookModel Book
	bookModel.Title = b.Title
	bookModel.Author = b.Author

	log.Println("Update Book...")

	if &b == nil {
		return nil, status.Error(codes.InvalidArgument, "can't be empty")
	}

	if err := configs.DB.Where("ID = ?", b.ID).First(b).Error; err != nil {
		return nil, err
	}

	if err := configs.DB.Model(b).Updates(bookModel).Error; err != nil {
		return nil, err
	}

	return b, nil
}

func (s *Server) DeleteBook(ctx context.Context, b *ID) (*ID, error) {
	var bookModel Book

	log.Println("Delete Book...")

	if err := configs.DB.Where("id = ? ", b.ID).First(&bookModel).Error; err != nil {
		return nil, err
	}

	configs.DB.Delete(bookModel)

	return b, nil
}

func (s *Server) ReadPublisher(ctx context.Context, p *ID) (*ListPublishers, error) {
	var book models.Book

	if err := configs.DB.Preload("Publishers").Where("id = ?", p.ID).First(&book).Error; err != nil {
		return nil, err
	}

	a := []*Publisher{}
	for i := 0; i < len(book.Publishers); i++ {
		b := Publisher{
			ID:     book.Publishers[i].ID,
			BookID: book.Publishers[i].BookID,
			Name:   book.Publishers[i].Name,
		}
		a = append(a, &b)
	}
	res := ListPublishers{ListPublisher: a}

	return &res, nil
}
