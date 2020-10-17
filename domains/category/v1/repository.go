package categoryv1

import (
	"encoding/json"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *CategoryV1) CreateCategory(request *models.Category) (*models.Category, error) {

	request.CreateTimestamp()

	result, err := o.orm.InsertInto("categorys", request)
	if err != nil {
		return nil, err
	}

	return result.(*models.Category), nil
}

func (o *CategoryV1) GetCategoryByID(id int64) (data *models.Category, err error) {
	data = &models.Category{}

	conn := *o.db
	if o.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = conn.Get(data, "select * from production.categorys where id=$1", id)
	if err != nil {
		return nil, err
	}

	o.log.Debug().Msgf("user %+v", data)

	return
}

func (d *CategoryV1) SearchCategorysByName(value *models.Search) (data *ArrayOfCategoryData, err error) {
	conn := *d.db
	if d.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	elementType := int64(value.ExtFilter["element_type"].(float64))

	rows, err := conn.Queryx(conn.Rebind("select * from production.categorys where element_type=$1 and description ilike $2"), elementType, value.Value+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data = &ArrayOfCategoryData{}

	for rows.Next() {
		var item models.Category

		err = rows.StructScan(&item)
		if err != nil {
			return nil, err
		}

		*data = append(*data, item)
	}

	return data, nil
}

func mergeCategoryData(oldData *models.Category, patch *[]byte) (newData *models.Category, err error) {
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

func (u *CategoryV1) UpdateCategoryByID(id int64, patch *[]byte) (writeData *models.Category, err error) {
	data, err := u.GetCategoryByID(id)
	if err != nil {
		return
	}

	writeData, err = mergeCategoryData(data, patch)
	if err != nil {
		return
	}

	if u.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	_, err = u.orm.Update("categorys", writeData)
	if err != nil {
		return nil, err
	}

	return writeData, err
}

func (u *CategoryV1) SoftDeleteCategoryByID(id int64) (err error) {
	data, err := u.GetCategoryByID(id)
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

	_, err = u.orm.Update("categorys", data)

	return
}

func (u *CategoryV1) HardDeleteCategoryByID(id int64) (err error) {
	conn := *u.db
	if conn == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = conn.Exec(conn.Rebind("DELETE FROM production.categorys WHERE id=$1"), id)

	if err != nil {
		return err
	}

	return nil
}
