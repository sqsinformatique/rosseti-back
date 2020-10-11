package userv1

import (
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (b *UserV1) CreateUser(request *models.User) (*models.User, error) {
	data := *request

	data.CreateTimestamp()

	result, err := b.orm.InsertInto("users", &data)
	if err != nil {
		return nil, err
	}

	return result.(*models.User), nil
}

func (b *UserV1) GetUserByID(id int) (data *models.User, err error) {
	data = &models.User{}

	if b.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = b.db.Get(data, "select * from production.users where id=$1", id)
	if err != nil {
		return nil, err
	}

	b.log.Debug().Msgf("user %+v", data)

	return
}
