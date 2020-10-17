package reviewv1

import (
	"encoding/json"
	"time"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
)

func (o *ReviewV1) CreateReview(request *models.Review) (*models.Review, error) {

	request.CreateTimestamp()

	result, err := o.orm.InsertInto("reviews", request)
	if err != nil {
		return nil, err
	}

	return result.(*models.Review), nil
}

func (o *ReviewV1) GetReviewByID(id int64) (data *models.Review, err error) {
	data = &models.Review{}

	conn := *o.db
	if o.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = conn.Get(data, "select * from production.reviews where id=$1", id)
	if err != nil {
		return nil, err
	}

	o.log.Debug().Msgf("user %+v", data)

	return
}

func mergeReviewData(oldData *models.Review, patch *[]byte) (newData *models.Review, err error) {
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

func (u *ReviewV1) UpdateReviewByID(id int64, patch *[]byte) (writeData *models.Review, err error) {
	data, err := u.GetReviewByID(id)
	if err != nil {
		return
	}

	writeData, err = mergeReviewData(data, patch)
	if err != nil {
		return
	}

	if u.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	_, err = u.orm.Update("reviews", writeData)
	if err != nil {
		return nil, err
	}

	return writeData, err
}

func (u *ReviewV1) SoftDeleteReviewByID(id int64) (err error) {
	data, err := u.GetReviewByID(id)
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

	_, err = u.orm.Update("reviews", data)

	return
}

func (u *ReviewV1) HardDeleteReviewByID(id int64) (err error) {
	conn := *u.db
	if conn == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = conn.Exec(conn.Rebind("DELETE FROM production.reviews WHERE id=$1"), id)

	if err != nil {
		return err
	}

	return nil
}
