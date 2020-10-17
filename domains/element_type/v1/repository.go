package elementtypev1

import (
	"encoding/json"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *ElementTypeV1) CreateElementType(request *models.ElementType) (*models.ElementType, error) {

	request.CreateTimestamp()

	result, err := o.orm.InsertInto("element_types", request)
	if err != nil {
		return nil, err
	}

	return result.(*models.ElementType), nil
}

func (o *ElementTypeV1) GetElementTypeByID(id int64) (data *models.ElementType, err error) {
	data = &models.ElementType{}

	conn := *o.db
	if o.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = conn.Get(data, "select * from production.element_types where id=$1", id)
	if err != nil {
		return nil, err
	}

	o.log.Debug().Msgf("user %+v", data)

	return
}

func mergeElementTypeData(oldData *models.ElementType, patch *[]byte) (newData *models.ElementType, err error) {
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

func (u *ElementTypeV1) UpdateElementTypeByID(id int64, patch *[]byte) (writeData *models.ElementType, err error) {
	data, err := u.GetElementTypeByID(id)
	if err != nil {
		return
	}

	writeData, err = mergeElementTypeData(data, patch)
	if err != nil {
		return
	}

	if u.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	_, err = u.orm.Update("element_types", writeData)
	if err != nil {
		return nil, err
	}

	return writeData, err
}

func (u *ElementTypeV1) SoftDeleteElementTypeByID(id int64) (err error) {
	data, err := u.GetElementTypeByID(id)
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

	_, err = u.orm.Update("element_types", data)

	return
}

func (u *ElementTypeV1) HardDeleteElementTypeByID(id int64) (err error) {
	conn := *u.db
	if conn == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = conn.Exec(conn.Rebind("DELETE FROM production.element_types WHERE id=$1"), id)

	if err != nil {
		return err
	}

	return nil
}
