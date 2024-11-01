package service

import (
	"errors"

	"github.com/betterreads/internal/domains/books/models"
	er "github.com/betterreads/internal/pkg/errors"
	"github.com/google/uuid"
)

var (
	

	ErrRatingNotFound = errors.New("rating not found")
	ErrBookNotFound   = errors.New("book not found")
    ErrPictureNotFound = errors.New("picture not found")
    ErrRatingAlreadyExists = errors.New("rating already exists")
    ErrReviewAlreadyExists = errors.New("review already exists")
    ErrReviewNotFound = errors.New("review not found")
	ErrAuthorNotFound = errors.New("author not found")
    ErrUserNotFound = errors.New("user not found")

    ErrGenreRequired= er.ErrorParam{
        Name:   "genre",
        Reason: "genre is required",
    }

    ErrGenreNotFound = er.ErrorParam{
        Name:   "genre",
        Reason: "genre not in available genres",
    }

	ErrRatingAmount = er.ErrorParam{
		Name:   "rating",
		Reason: "rating must be between 1 and 5",
	}

)

type BooksService interface {
    PublishBook(req *models.NewBookRequest, author uuid.UUID) (*models.BookResponse, error)
    GetBookInfo(bookId uuid.UUID, userId uuid.UUID) (*models.BookResponseWithReview, error)
    GetBooksOfAuthor(authorId uuid.UUID, userId uuid.UUID) ([]*models.BookResponseWithReview, error)
    SearchBooksByName(name string, userId uuid.UUID) ([]*models.BookResponseWithReview, error)
    GetBookPicture(id uuid.UUID) ([]byte, error)
    GetBooksInfo(userId uuid.UUID) ([]*models.BookResponseWithReview, error)
    RateBook(bookId uuid.UUID, userId uuid.UUID, rateAmount int) (*models.Rating, error)
    UpdateRating(bookId uuid.UUID, userId uuid.UUID, rateAmount int) (error)
    GetBookReviews(bookId uuid.UUID) ([]*models.Review, error)
    GetAllReviewsOfUser(userId uuid.UUID) ([]*models.Review, error)
    AddReview(bookId uuid.UUID, userId uuid.UUID, review models.NewReviewRequest) error
}

