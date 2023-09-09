package user

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type validateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type validateUserResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

func makeValidateUserEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(validateUserRequest)
		if !ok {
			return nil, errors.New("Invalid Request")
		}
		token, err := svc.ValidateUser(ctx, req.Email, req.Password)
		if err != nil {
			return validateUserResponse{Token: "", Error: err.Error()}, err
		}
		return validateUserResponse{token, ""}, nil

	}
}
