package helpers

import (
	"errors"
	"net/http"
	"regexp"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func BindQuery(ctx *gin.Context, query *structs.Query) error {
	err := ctx.ShouldBindQuery(query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return errors.New("invalid request")
	}
	if query.Page == 0 {
		query.Page = 1
	}
	if query.PerPage == 0 {
		query.PerPage = 10
	}

	return nil
}

func BindModelQuery[T structs.UserQuery | structs.BookQuery | structs.VisitationQuery | structs.BookReadsQuery | structs.BorrowedQuery](ctx *gin.Context, query *T) error {
	err := ctx.ShouldBindQuery(query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return errors.New("invalid request")
	}
	return nil
}

func GetOffset(query structs.Query) int {
	return (query.Page - 1) * query.PerPage
}

func GenerateMeta(count int64, query structs.Query) structs.Meta {
	var nextPage int
	if count > int64(query.PerPage) {
		nextPage = query.Page + 1
	} else {
		nextPage = 0
	}
	nextPageAvailable := count > int64(query.PerPage)
	meta := structs.Meta{
		TotalRecords:      count,
		Page:              query.Page,
		PerPage:           query.PerPage,
		NextPage:          nextPage,
		NextPageAvailable: nextPageAvailable,
	}
	return meta
}

func EncryptPasswordFromPlaintext(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CompareHashedPasswordWithPlaintext(hashed, plaintext string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plaintext))
	return err == nil
}

func ValidatePasswordPlaintext(password string) error {
	if password == "" {
		return errors.New("password must be provider")
	}
	if len(password) < 8 {
		return errors.New("password characters should be at least 8")
	}
	if len(password) > 40 {
		return errors.New("password can not have more than 20 characters")
	}
	return nil
}

func ValidateEmail (email string) error {
	if email == "" {
		return errors.New("email must be provided")
	}
	//check that email is valid with regex
	if !Matches(email, EmailRX) {
		return errors.New("email is not valid")
	}
	return nil
}

func CheckThatUserHasValidValues(user models.User) error {
	if user.Name == "" {
		return errors.New("name must be provided")
	}
	if err := ValidateEmail(user.Email); err != nil {
		return err
	}

	if err := ValidatePasswordPlaintext(user.Password); err != nil {
		return	err
	}
	return nil
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
