package objectv1

import (
	"encoding/json"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *ObjectV1) CreateObject(request *models.Object) (*models.Object, error) {

	request.CreateTimestamp()

	result, err := o.orm.InsertInto("objects", request)
	if err != nil {
		return nil, err
	}

	return result.(*models.Object), nil
}

func (o *ObjectV1) GetObjectByID(id int64) (data *models.Object, err error) {
	data = &models.Object{}

	conn := *o.db
	if conn == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = conn.Get(data, "select * from production.objects where id=$1", id)
	if err != nil {
		return nil, err
	}

	o.log.Debug().Msgf("user %+v", data)

	return
}

func mergeObjectData(oldData *models.Object, patch *[]byte) (newData *models.Object, err error) {
	id := oldData.ID

	original, err := json.Marshal(oldData)
	if err != nil {
		return
	}

	merged, err := jsonpatch.MergePatch(original, *patch)
	if err != nil {
		return
	}

	err = json.Unmarshal(merged, &newData)
	if err != nil {
		return
	}

	// Protect ID from changes
	newData.ID = id

	newData.UpdatedAt.Time = time.Now()
	newData.UpdatedAt.Valid = true

	return newData, nil
}

func (u *ObjectV1) UpdateObjectByID(id int64, patch *[]byte) (writeData *models.Object, err error) {
	data, err := u.GetObjectByID(id)
	if err != nil {
		return
	}

	writeData, err = mergeObjectData(data, patch)
	if err != nil {
		return
	}

	if u.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	_, err = u.orm.Update("objects", writeData)
	if err != nil {
		return nil, err
	}

	return writeData, err
}

func (u *ObjectV1) SoftDeleteObjectByID(id int64) (err error) {
	data, err := u.GetObjectByID(id)
	if err != nil {
		return
	}

	if data.DeletedAt.Valid {
		return
	}

	data.DeletedAt.Time = time.Now()
	data.DeletedAt.Valid = true
	data.UpdatedAt.Time = time.Now()
	data.UpdatedAt.Valid = true

	if u.db == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = u.orm.Update("objects", data)

	return
}

func (u *ObjectV1) HardDeleteObjectByID(id int64) (err error) {
	conn := *u.db
	if conn == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = conn.Exec(conn.Rebind("DELETE FROM production.objects WHERE id=$1"), id)

	if err != nil {
		return err
	}

	return nil
}
