package elementequipmentv1

import (
	"encoding/json"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *ElementEquipmentV1) CreateElementEquipment(request *models.ElementEquipment) (*models.ElementEquipment, error) {

	request.CreateTimestamp()

	result, err := o.orm.InsertInto("equipments", request)
	if err != nil {
		return nil, err
	}

	return result.(*models.ElementEquipment), nil
}

func (o *ElementEquipmentV1) GetElementEquipmentByID(id int64) (data *models.ElementEquipment, err error) {
	data = &models.ElementEquipment{}

	conn := *o.db
	if o.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = conn.Get(data, "select * from production.equipments where id=$1", id)
	if err != nil {
		return nil, err
	}

	o.log.Debug().Msgf("user %+v", data)

	return
}

func (d *ElementEquipmentV1) SearchElementEquipmentsByName(value *models.Search) (data *ArrayOfElementEquipmentData, err error) {
	conn := *d.db
	if d.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	elementType := int64(value.ExtFilter["element_type"].(float64))

	rows, err := conn.Queryx(conn.Rebind("select * from production.equipments where element_type=$1 and description ilike $2"), elementType, value.Value+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data = &ArrayOfElementEquipmentData{}

	for rows.Next() {
		var item models.ElementEquipment

		err = rows.StructScan(&item)
		if err != nil {
			return nil, err
		}

		*data = append(*data, item)
	}

	return data, nil
}

func mergeElementEquipmentData(oldData *models.ElementEquipment, patch *[]byte) (newData *models.ElementEquipment, err error) {
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

func (u *ElementEquipmentV1) UpdateElementEquipmentByID(id int64, patch *[]byte) (writeData *models.ElementEquipment, err error) {
	data, err := u.GetElementEquipmentByID(id)
	if err != nil {
		return
	}

	writeData, err = mergeElementEquipmentData(data, patch)
	if err != nil {
		return
	}

	if u.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	_, err = u.orm.Update("equipments", writeData)
	if err != nil {
		return nil, err
	}

	return writeData, err
}

func (u *ElementEquipmentV1) SoftDeleteElementEquipmentByID(id int64) (err error) {
	data, err := u.GetElementEquipmentByID(id)
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

	_, err = u.orm.Update("element_equipments", data)

	return
}

func (u *ElementEquipmentV1) HardDeleteElementEquipmentByID(id int64) (err error) {
	conn := *u.db
	if conn == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = conn.Exec(conn.Rebind("DELETE FROM production.equipments WHERE id=$1"), id)

	if err != nil {
		return err
	}

	return nil
}
