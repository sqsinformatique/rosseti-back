package orderv1

import (
	"encoding/json"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *OrderV1) CreateOrder(request *models.Order) (*models.Order, error) {

	request.CreateTimestamp()

	result, err := o.orm.InsertInto("orders", request)
	if err != nil {
		return nil, err
	}

	return result.(*models.Order), nil
}

func (o *OrderV1) GetOrderByID(id int64) (data *models.Order, err error) {
	data = &models.Order{}

	conn := *o.db
	if conn == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = conn.Get(data, "select * from production.orders where id=$1", id)
	if err != nil {
		return nil, err
	}

	o.log.Debug().Msgf("user %+v", data)
	staff, err := o.profileV1.GetProfileByID(int64(data.StaffID))
	if err != nil {
		return nil, err
	}
	o.log.Debug().Msgf("staff %+v", staff)
	data.StaffDesc = staff

	superviser, err := o.profileV1.GetProfileByID(int64(data.SuperviserID))
	if err != nil {
		return nil, err
	}

	data.SuperviserDesc = superviser

	object, err := o.objectV1.GetObjectByID(int64(data.ObjectID))
	if err != nil {
		return nil, err
	}

	data.ObjectDesc = object

	var techtasks []*models.TechTask
	for _, v := range data.TechTasks.Map {
		for _, w := range v.([]interface{}) {
			z := w.(map[string]interface{})
			task, err := o.techtaskV1.GetTechTaskByID(int64(z["task_id"].(float64)))
			if err != nil {
				return nil, err
			}

			techtasks = append(techtasks, task)
		}
	}

	data.TechTasksDesc = techtasks

	return
}

func (o *OrderV1) GetOrdersByUserID(id int64) (data *ArrayOfOrderData, err error) {
	conn := *o.db
	if conn == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	rows, err := conn.Queryx(conn.Rebind("select * from production.orders where staff_id=$1"), id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data = &ArrayOfOrderData{}

	for rows.Next() {
		var item models.Order

		err = rows.StructScan(&item)
		if err != nil {
			return nil, err
		}

		staff, err := o.profileV1.GetProfileByID(int64(item.StaffID))
		if err != nil {
			return nil, err
		}
		o.log.Debug().Msgf("staff %+v", staff)
		item.StaffDesc = staff

		superviser, err := o.profileV1.GetProfileByID(int64(item.SuperviserID))
		if err != nil {
			return nil, err
		}

		item.SuperviserDesc = superviser

		object, err := o.objectV1.GetObjectByID(int64(item.ObjectID))
		if err != nil {
			return nil, err
		}

		item.ObjectDesc = object

		var techtasks []*models.TechTask
		for _, v := range item.TechTasks.Map {
			for _, w := range v.([]interface{}) {
				z := w.(map[string]interface{})
				task, err := o.techtaskV1.GetTechTaskByID(int64(z["task_id"].(float64)))
				if err != nil {
					return nil, err
				}

				techtasks = append(techtasks, task)
			}
		}

		item.TechTasksDesc = techtasks

		*data = append(*data, item)
	}

	return data, err
}

func mergeOrderData(oldData *models.Order, patch *[]byte) (newData *models.Order, err error) {
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

func (u *OrderV1) UpdateOrderByID(id int64, patch *[]byte) (writeData *models.Order, err error) {
	data, err := u.GetOrderByID(id)
	if err != nil {
		return
	}

	writeData, err = mergeOrderData(data, patch)
	if err != nil {
		return
	}

	if u.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	_, err = u.orm.Update("orders", writeData)
	if err != nil {
		return nil, err
	}

	return writeData, err
}

func (u *OrderV1) SoftDeleteOrderByID(id int64) (err error) {
	data, err := u.GetOrderByID(id)
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

	_, err = u.orm.Update("orders", data)

	return
}

func (u *OrderV1) HardDeleteOrderByID(id int64) (err error) {
	conn := *u.db
	if u.db == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = conn.Exec(conn.Rebind("DELETE FROM production.orders WHERE id=$1"), id)

	if err != nil {
		return err
	}

	return nil
}
