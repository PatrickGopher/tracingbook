package dtos

import (
	"net/http"
	"tracingbook/models"
)

func CreateHomeResponse(request *http.Request, tags []models.Book) map[string]interface{} {
	return CreateSuccessDto(map[string]interface{}{
		"books": CreateBookListDto(request, tags),
	})
}
