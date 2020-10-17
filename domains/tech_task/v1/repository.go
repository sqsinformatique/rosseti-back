package techtaskv1

import (
	"encoding/json"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *TechTaskV1) CreateTechTask(request *models.TechTask) (*models.TechTask, error) {

	request.CreateTimestamp()

	result, err := o.orm.InsertInto("tech_tasks", request)
	if err != nil {
		return nil, err
	}

	return result.(*models.TechTask), nil
}

func (o *TechTaskV1) GetTechTaskByID(id int64) (data *models.TechTask, err error) {
	data = &models.TechTask{}

	conn := *o.db
	if o.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = conn.Get(data, "select * from production.tech_tasks where id=$1", id)
	if err != nil {
		return nil, err
	}

	o.log.Debug().Msgf("user %+v", data)

	return
}

func mergeTechTaskData(oldData *models.TechTask, patch *[]byte) (newData *models.TechTask, err error) {
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

func (u *TechTaskV1) UpdateTechTaskByID(id int64, patch *[]byte) (writeData *models.TechTask, err error) {
	data, err := u.GetTechTaskByID(id)
	if err != nil {
		return
	}

	writeData, err = mergeTechTaskData(data, patch)
	if err != nil {
		return
	}

	if u.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	_, err = u.orm.Update("tech_tasks", writeData)
	if err != nil {
		return nil, err
	}

	return writeData, err
}

func (u *TechTaskV1) SoftDeleteTechTaskByID(id int64) (err error) {
	data, err := u.GetTechTaskByID(id)
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

	_, err = u.orm.Update("tech_tasks", data)

	return
}

func (u *TechTaskV1) HardDeleteTechTaskByID(id int64) (err error) {
	conn := *u.db
	if conn == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = conn.Exec(conn.Rebind("DELETE FROM production.tech_tasks WHERE id=$1"), id)

	if err != nil {
		return err
	}

	return nil
}
