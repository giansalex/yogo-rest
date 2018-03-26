package rest_api

import (
	"context"
)

type Service interface {
	Get(ctx context.Context, obj rest_apiRequest) (rest_apiResponse, error)
}

type Repository interface {
	Get(ID string) (rest_apiResponse, error)
}
