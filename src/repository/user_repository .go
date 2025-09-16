package repository

import (
	"github.com/fastgox/fastgox-api-starter/src/models/entity"
)

// UserRepository 用户仓储
type UserRepository struct {
	BaseRepository[entity.User]
}

var UserRepo = &UserRepository{BaseRepository: *GetRepository[entity.User]()}
