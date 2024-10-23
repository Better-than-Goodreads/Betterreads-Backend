package service

import (
	"errors"

	"github.com/betterreads/internal/domains/books/models"
	"github.com/betterreads/internal/domains/books/repository"
	"github.com/betterreads/internal/domains/books/utils"
	er "github.com/betterreads/internal/pkg/errors"
	"github.com/google/uuid"
)

var (
	ErrGenreNotFound  = errors.New("genre not found")
	ErrRatingNotFound = errors.New("rating not found")

	ErrRatingAmount = er.ErrorParam{
		Name: "rating",
		Reason: "rating must be between 1 and 5",
	}
)

type BooksService struct {
	booksRepository repository.BooksDatabase
}

func NewBooksService(booksRepository repository.BooksDatabase) *BooksService {
	return &BooksService{booksRepository: booksRepository}
}

func (bs *BooksService) PublishBook(req *models.NewBookRequest, author uuid.UUID) (*models.BookResponse, error) {
	if len(req.Genres) == 0 {
		return nil, errors.New("at least one genre is required")
	}

	book, err := bs.booksRepository.SaveBook(req, author)
	if err != nil {
		return nil, err
	}

    bookRes, err := bs.addAuthor(book, book.Author)
    if err != nil {
        return nil, err
    }

	return bookRes, nil
}

func (bs *BooksService) GetBook(id uuid.UUID) (*models.BookResponse, error) {
	book, err := bs.booksRepository.GetBookById(id)
	if err != nil {
		return nil, err
	}
    
    bookRes, err := bs.addAuthor(book, book.Author)
    if err != nil {
        return nil, err
    }


    return bookRes, nil
}

func (bs *BooksService) GetBooks() ([]*models.BookResponse, error) {
	books, err := bs.booksRepository.GetBooks()
	if err != nil {
		return nil, err
	}
    
    booksResponses := []*models.BookResponse{}
    for _, book := range books {
        bookResponse, err := bs.addAuthor(book, book.Author)
        booksResponses = append(booksResponses, bookResponse)
        if err != nil {
            return nil, err
        }
    }
	return booksResponses, nil
}

func (bs *BooksService) RateBook(bookId uuid.UUID, userId uuid.UUID, rateAmount int) error {

	if rateAmount < 1 || rateAmount > 5 {
		return	ErrRatingAmount 
	}

	err := bs.booksRepository.RateBook(bookId, userId, rateAmount)
	if err != nil {
		return err
	}
	return nil
}

func (bs *BooksService) DeleteRating(bookId uuid.UUID, userId uuid.UUID) error {
	err := bs.booksRepository.DeleteRating(bookId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (bs *BooksService) GetRatingUser(bookId uuid.UUID, userId uuid.UUID) (*models.RatingResponse, error) {
	rating, err := bs.booksRepository.GetRatingUser(bookId, userId)
	if err != nil {
		if errors.Is(err, repository.ErrRatingNotFound) {
			return nil, ErrRatingNotFound
		}
		return nil, err
	}
	ratingResponse := utils.MapRatingToRatingResponse(rating)
	return ratingResponse, nil
}


func (bs *BooksService) addAuthor(book *models.Book, author uuid.UUID) (*models.BookResponse, error) {
    author_name, err := bs.booksRepository.GetAuthorName(author)
    if err != nil {
        return nil, err
    }
    bookRes := utils.MapBookToBookResponse(book, author_name)
    return  bookRes , nil
}
