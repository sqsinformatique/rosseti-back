package actsdetailv1

import (
	"encoding/json"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *ActsDetailV1) CreateActsDetail(request *models.ActsDetail) (*models.ActsDetail, error) {

	request.CreateTimestamp()

	result, err := o.orm.InsertInto("acts_details", request)
	if err != nil {
		return nil, err
	}

	return result.(*models.ActsDetail), nil
}

func (o *ActsDetailV1) GetAllActsDetailByID(id, objectID int64) (data *ArrayOfActsDetailData, err error) {
	conn := *o.db
	if o.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	rows, err := conn.Queryx(conn.Rebind("select * from production.acts_details where act_id=$1"), id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data = &ArrayOfActsDetailData{}

	for rows.Next() {
		var item models.ActsDetail

		err = rows.StructScan(&item)
		if err != nil {
			return nil, err
		}

		var defects []*models.Defect
		for _, v := range item.Defects.Map {
			for _, w := range v.([]interface{}) {
				z := w.(map[string]interface{})
				task, err := o.defectV1.GetDefectByID(int64(z["defect_id"].(float64)))
				if err != nil {
					return nil, err
				}

				defects = append(defects, task)
			}
		}

		item.DefectsDef = defects

		elementDesc, err := o.objectsdetailV1.GetObjectsDetailByID(objectID, int64(item.ElementID))
		if err != nil {
			return nil, err
		}

		item.ElementDesc = elementDesc

		*data = append(*data, item)
	}

	o.log.Debug().Msgf("user %+v", data)

	return
}

func (o *ActsDetailV1) GetActsDetailByID(id int64) (data *models.ActsDetail, err error) {
	data = &models.ActsDetail{}

	conn := *o.db
	if o.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = conn.Get(data, "select * from production.acts_details where act_id=$1", id)
	if err != nil {
		return nil, err
	}

	o.log.Debug().Msgf("user %+v", data)

	return
}

func mergeActsDetailData(oldData *models.ActsDetail, patch *[]byte) (newData *models.ActsDetail, err error) {
	actid := oldData.ActID

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
	newData.ActID = actid

	newData.UpdatedAt.Time = time.Now()
	newData.UpdatedAt.Valid = true

	return newData, nil
}

func (u *ActsDetailV1) UpdateActsDetailByID(id int64, patch *[]byte) (writeData *models.ActsDetail, err error) {
	data, err := u.GetActsDetailByID(id)
	if err != nil {
		return
	}

	writeData, err = mergeActsDetailData(data, patch)
	if err != nil {
		return
	}

	if u.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	_, err = u.orm.Update("acts_details", writeData)
	if err != nil {
		return nil, err
	}

	return writeData, err
}

func (u *ActsDetailV1) SoftDeleteActsDetailByID(id int64) (err error) {
	data, err := u.GetActsDetailByID(id)
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

	_, err = u.orm.Update("acts_details", data)

	return
}

func (u *ActsDetailV1) HardDeleteActsDetailByID(id int64) (err error) {
	conn := *u.db
	if conn == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = conn.Exec(conn.Rebind("DELETE FROM production.acts_details WHERE id=$1"), id)

	if err != nil {
		return err
	}

	return nil
}
