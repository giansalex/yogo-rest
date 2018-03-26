package rest_api

import "context"

//Service...
type service struct {
	Repository Repository
}

//NewService ...
func NewService(r Repository) Service {
	return service{Repository: r}
}

//Get ...
func (s service) Get(ctx context.Context, obj rest_apiRequest) (rest_apiResponse, error) {
	return s.Repository.Get(obj.ID)
}
