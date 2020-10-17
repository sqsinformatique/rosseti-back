package defectv1

import (
	"encoding/json"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *DefectV1) CreateDefect(request *models.Defect) (*models.Defect, error) {

	request.CreateTimestamp()

	result, err := o.orm.InsertInto("defects", request)
	if err != nil {
		return nil, err
	}

	return result.(*models.Defect), nil
}

func (o *DefectV1) GetDefectByID(id int64) (data *models.Defect, err error) {
	data = &models.Defect{}

	conn := *o.db
	if o.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = conn.Get(data, "select * from production.defects where id=$1", id)
	if err != nil {
		return nil, err
	}

	o.log.Debug().Msgf("user %+v", data)

	return
}

func (d *DefectV1) SearchDefectsByName(value *models.Search) (data *ArrayOfDefectData, err error) {
	conn := *d.db
	if d.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	elementType := int64(value.ExtFilter["element_type"].(float64))

	rows, err := conn.Queryx(conn.Rebind("select * from production.defects where element_type=$1 and description ilike $2"), elementType, value.Value+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data = &ArrayOfDefectData{}

	for rows.Next() {
		var item models.Defect

		err = rows.StructScan(&item)
		if err != nil {
			return nil, err
		}

		*data = append(*data, item)
	}

	return data, nil
}

func mergeDefectData(oldData *models.Defect, patch *[]byte) (newData *models.Defect, err error) {
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

func (u *DefectV1) UpdateDefectByID(id int64, patch *[]byte) (writeData *models.Defect, err error) {
	data, err := u.GetDefectByID(id)
	if err != nil {
		return
	}

	writeData, err = mergeDefectData(data, patch)
	if err != nil {
		return
	}

	if u.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	_, err = u.orm.Update("defects", writeData)
	if err != nil {
		return nil, err
	}

	return writeData, err
}

func (u *DefectV1) SoftDeleteDefectByID(id int64) (err error) {
	data, err := u.GetDefectByID(id)
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

	_, err = u.orm.Update("defects", data)

	return
}

func (u *DefectV1) HardDeleteDefectByID(id int64) (err error) {
	conn := *u.db
	if conn == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = conn.Exec(conn.Rebind("DELETE FROM production.defects WHERE id=$1"), id)

	if err != nil {
		return err
	}

	return nil
}
