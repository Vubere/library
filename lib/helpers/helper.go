package helpers

import (
	"errors"
	"net/http"
	"victorubere/library/lib/structs"

	"github.com/gin-gonic/gin"
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
func BindModelQuery[T structs.UserQuery | structs.BookQuery | structs.VisitationQuery | structs.BookReadQuery | structs.BorrowedQuery](ctx *gin.Context, query *T) error {
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
