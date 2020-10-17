package objectsdetailv1

import (
	"encoding/json"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *ObjectsDetailV1) CreateObjectsDetail(request *models.ObjectsDetail) (*models.ObjectsDetail, error) {

	request.CreateTimestamp()

	result, err := o.orm.InsertInto("objects_details", request)
	if err != nil {
		return nil, err
	}

	return result.(*models.ObjectsDetail), nil
}

func (o *ObjectsDetailV1) GetObjectsDetailByID(objectID, elementID int64) (data *models.ObjectsDetail, err error) {
	data = &models.ObjectsDetail{}

	conn := *o.db
	if o.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = conn.Get(data, "select * from production.objects_details where object_id=$1 and element_id=$2", objectID, elementID)
	if err != nil {
		return nil, err
	}

	o.log.Debug().Msgf("user %+v", data)

	return
}

func (o *ObjectsDetailV1) SearchObjectsDetailByName(value *models.Search) (data *ArrayOfObjectsDetailData, err error) {
	conn := *o.db
	if o.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	objectID := int64(value.ExtFilter["object_id"].(float64))

	rows, err := conn.Queryx(conn.Rebind("select * from production.objects_details where object_id=$1 and element_name ilike $2"), objectID, value.Value+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data = &ArrayOfObjectsDetailData{}

	for rows.Next() {
		var item models.ObjectsDetail

		err = rows.StructScan(&item)
		if err != nil {
			return nil, err
		}

		*data = append(*data, item)
	}

	return data, nil
}

func mergeObjectsDetailData(oldData *models.ObjectsDetail, patch *[]byte) (newData *models.ObjectsDetail, err error) {
	objectID := oldData.ObjectID
	elementID := oldData.ElementID

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
	newData.ObjectID = objectID
	newData.ElementID = elementID

	newData.UpdatedAt.Time = time.Now()
	newData.UpdatedAt.Valid = true

	return newData, nil
}

func (u *ObjectsDetailV1) UpdateObjectsDetailByID(objectID, elementID int64, patch *[]byte) (writeData *models.ObjectsDetail, err error) {
	data, err := u.GetObjectsDetailByID(objectID, elementID)
	if err != nil {
		return
	}

	writeData, err = mergeObjectsDetailData(data, patch)
	if err != nil {
		return
	}

	if u.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	_, err = u.orm.Update("objects_details", writeData)
	if err != nil {
		return nil, err
	}

	return writeData, err
}

func (u *ObjectsDetailV1) SoftDeleteObjectsDetailByID(objectID, elementID int64) (err error) {
	data, err := u.GetObjectsDetailByID(objectID, elementID)
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

	_, err = u.orm.Update("objects_details", data)

	return
}

func (u *ObjectsDetailV1) HardDeleteObjectsDetailByID(objectID, elementID int64) (err error) {
	conn := *u.db
	if conn == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = conn.Exec(conn.Rebind("DELETE FROM production.objects_details WHERE object_id=$1 and element_id=$2"), objectID, elementID)

	if err != nil {
		return err
	}

	return nil
}
