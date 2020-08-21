package vo

import "go-web-base/core/entity"

type LoginResult struct {
	User entity.User
	Token string
	ExpirationIn int

}
