package rest_api

import (
	"context"

	tracing "github.com/ricardo-ch/go-tracing"
)

type rest_apiTracing struct {
	next Service
}

// NewTracing ...
func NewTracing(s Service) Service {
	return rest_apiTracing{
		next: s,
	}
}

// Get ...
func (s rest_apiTracing) Get(ctx context.Context, obj rest_apiRequest) (response rest_apiResponse, err error) {
	span, ctx := tracing.CreateSpan(ctx, "rest_api.service::Get", &map[string]interface{}{"id": obj.ID})
	defer func() {
		if err != nil {
			tracing.SetSpanError(span, err)
		}
		span.Finish()
	}()

	return s.next.Get(ctx, obj)
}
