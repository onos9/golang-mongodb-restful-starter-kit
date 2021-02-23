package auth

import (
	"context"

	model "golang-mongodb-restful-starter-kit/app/models"
	"golang-mongodb-restful-starter-kit/config"

	repository "golang-mongodb-restful-starter-kit/app/repositories/user"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AuthService interface {
	Create(context.Context, *model.User) error
	Login(context.Context, *model.Credential) (*model.User, error)
	IsUserAlreadyExists(context.Context, string) bool
}

type AuthServiceImp struct {
	db         *mgo.Session
	repository repository.UserRepository
	config     *config.Configuration
}

func New(db *mgo.Session, c *config.Configuration) AuthService {
	return &AuthServiceImp{db: db, config: c, repository: repository.New(db, c)}
}

func (service *AuthServiceImp) Create(ctx context.Context, user *model.User) error {

	return service.repository.Create(ctx, user)
}

func (service *AuthServiceImp) Login(ctx context.Context, credential *model.Credential) (*model.User, error) {
	query := bson.M{"email": credential.Email}
	user, err := service.repository.FindOne(ctx, query)
	if err != nil {
		return nil, err
	}

	if err = user.ComparePassword(credential.Password); err != nil {
		return nil, err
	}
	return user, nil

}

// IsUserAlreadyExists , checks if user already exists in DB
func (service *AuthServiceImp) IsUserAlreadyExists(ctx context.Context, email string) bool {

	return service.repository.IsUserAlreadyExists(ctx, email)

}
