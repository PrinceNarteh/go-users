package responses

import "github.com/PrinceNarteh/go-users/models"

type UserResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Data       models.User `json:"data"`
}
