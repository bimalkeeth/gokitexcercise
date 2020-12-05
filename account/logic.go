package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	uui "github.com/satori/go.uuid"
)

type service struct {
	repository Repository
	logger log.Logger
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger:=log.With(s.logger,"method","create user")
	uuid:=uui.NewV4()
	id:=uuid.String()
	user :=User{
		ID: id,
		Email: email,
		Password: password,
	}
	if err:=s.repository.CreateUser(ctx,user);err!=nil{
		_ = level.Error(logger).Log("err", err)
		return "",err
	}
	_ = logger.Log("create user", id)

	return "Success",nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger:=log.With(s.logger,"method","get user")
	email,err:=s.repository.GetUser(ctx,id)
	if err!=nil{
		_ = level.Error(logger).Log("err", err)
		return "",err
	}
	_ = logger.Log("Get user", id)
	return email,nil
}

func NewService(rep Repository,logger log.Logger) Service{
	return &service{
		repository: rep,
		logger: logger,
	}
}