package profilev1

import (
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (b *ProfileV1) CreateProfile(request *models.Profile) (*models.Profile, error) {
	data := *request

	data.CreateTimestamp()

	result, err := b.orm.InsertInto("profiles", &data)
	if err != nil {
		return nil, err
	}

	return result.(*models.Profile), nil
}

func (b *ProfileV1) GetProfileByID(id int) (data *models.Profile, err error) {
	data = &models.Profile{}

	if b.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = b.db.Get(data, "select * from production.profiles where id=$1", id)
	if err != nil {
		return nil, err
	}

	b.log.Debug().Msgf("profile %+v", data)

	return
}
