package actv1

import (
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (b *ActV1) CreateAct(request *models.Act) (*models.Act, error) {
	data := *request

	data.CreateTimestamp()

	result, err := b.orm.InsertInto("acts", &data)
	if err != nil {
		return nil, err
	}

	return result.(*models.Act), nil
}

func (b *ActV1) GetActByID(id int) (data *models.Act, err error) {
	data = &models.Act{}

	if b.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = b.db.Get(data, "select * from production.acts where id=$1", id)
	if err != nil {
		return nil, err
	}

	b.log.Debug().Msgf("acts %+v", data)

	return
}
