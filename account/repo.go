package account

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
)

var RepoError =errors.New("unable to handle repo request")

type repo struct {
	db *sql.DB
	log.Logger
}

func (r *repo) CreateUser(ctx context.Context, user User) error {
	sqlStr:=`
             INSERT INTO users(id,email,password)
             VALUES($1,$2,$3)`
	if user.Email == "" || user.Password =="" {
		return RepoError
	}
	_,err:=r.db.ExecContext(ctx,sqlStr,user.ID,user.Email,user.Password)
	if err!=nil{
		return err
	}
	return nil
}

func (r *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string 
	err:=r.db.QueryRow("SELECT email FROM users WHERE  id=$1",id).Scan(&email)
	if err!=nil{
		return "", RepoError
	}
	return email, nil
}

func NewRepo(db *sql.DB,logger log.Logger) Repository{
	return &repo{
		db ,
		log.With(logger,"repo","sql"),
	}
}