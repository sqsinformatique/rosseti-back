package actv1

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
	"github.com/sqsinformatique/rosseti-back/types"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

var (
	ErrActIsFinished = errors.New("act is finished, can not be updated")
)

func (a *ActV1) CreateAct(request *models.Act) (*models.Act, error) {
	data := *request

	data.CreateTimestamp()

	if data.Finished {
		sign, err := a.profilev1.SignDataByID(int64(data.StaffID), &data)
		if err != nil {
			return nil, err
		}
		data.StaffSign = sign
	}

	result, err := a.orm.InsertInto("acts", &data)
	if err != nil {
		return nil, err
	}

	return result.(*models.Act), nil
}

func (a *ActV1) UpdateActByID(id string, request *models.Act) (*models.Act, error) {
	data := *request

	actID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	oldAct, err := a.GetActByID(id)
	if err != nil {
		return nil, err
	}

	if oldAct.Finished {
		return nil, ErrActIsFinished
	}

	data.ID = int(actID)
	data.UpdateTimestamp()

	result, err := a.orm.Update("acts", &data)
	if err != nil {
		return nil, err
	}

	if data.Finished {
		sign, err := a.profilev1.SignDataByID(int64(data.StaffID), &data)
		if err != nil {
			return nil, err
		}
		data.StaffSign = sign
	}

	return result.(*models.Act), nil
}

func (a *ActV1) GetActByID(id string) (data *models.Act, err error) {
	actID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	data = &models.Act{}

	conn := *a.db
	if conn == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = conn.Get(data, "select * from production.acts where id=$1", actID)
	if err != nil {
		return nil, err
	}

	a.log.Debug().Msgf("acts %+v", data)
	staff, err := a.profilev1.GetProfileByID(int64(data.StaffID))
	if err != nil {
		return nil, fmt.Errorf("failed gets staff: %w", err)
	}

	data.StaffDesc = staff

	superviser, err := a.profilev1.GetProfileByID(int64(data.SuperviserID))
	if err != nil {
		return nil, fmt.Errorf("failed gets superviser: %w", err)
	}

	data.SuperviserDesc = superviser

	object, err := a.objectV1.GetObjectByID(int64(data.ObjectID))
	if err != nil {
		return nil, fmt.Errorf("failed gets object: %w", err)
	}

	data.ObjectDecs = object

	return data, nil
}

func (a *ActV1) GetActsByDate(timeStart, timeEnd types.NullTime) (data *ArrayOfOActData, err error) {
	conn := *a.db
	if conn == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	rows, err := conn.Queryx(conn.Rebind("select * from production.acts where updated_at>=$1 and updated_at<=$2"), timeStart, timeEnd)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data = &ArrayOfOActData{}

	for rows.Next() {
		var item models.Act

		err = rows.StructScan(&item)
		if err != nil {
			return nil, err
		}

		*data = append(*data, item)
	}

	return data, nil
}

func (a *ActV1) GetActsByUserID(id string) (data *ArrayOfOActData, err error) {
	data = &ArrayOfOActData{}

	return data, nil
}

func writeToGridFile(fileName string, file multipart.File, gridFile *gridfs.UploadStream) (int, error) {
	reader := bufio.NewReader(file)
	defer func() { file.Close() }()
	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	fileSize := 0
	for {
		// read a chunk
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return 0, errors.New("could not read the input file")
		}
		if n == 0 {
			break
		}
		// write a chunk
		if size, err := gridFile.Write(buf[:n]); err != nil {
			return 0, errors.New("could not write to GridFs for " + fileName)
		} else {
			fileSize += size
		}
	}
	gridFile.Close()
	return fileSize, nil
}

func (a *ActV1) UpdateImagesList(actID, filename string) error {
	return nil
}

func (a *ActV1) CreateImages(actID string, multipartForm *multipart.Form) error {
	for _, fileHeaders := range multipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			mongoconn := *a.mongodb
			bucket, err := gridfs.NewBucket(
				mongoconn.Database(a.cfg.Mongo.ImageDB),
			)
			if err != nil {
				return err
			}

			gridFile, err := bucket.OpenUploadStream(
				actID + "_" + fileHeader.Filename, // this is the name of the file which will be saved in the database
			)
			if err != nil {
				return err
			}

			fileSize, err := writeToGridFile(fileHeader.Filename, file, gridFile)
			if err != nil {
				return err
			}

			a.log.Debug().Msgf("Write file to DB was successful. File size: %d \n", fileSize)
		}
	}

	return nil
}

func (a *ActV1) GetImage(actID, imageID string) (*bytes.Buffer, int64, error) {
	mongoconn := *a.mongodb
	bucket, err := gridfs.NewBucket(
		mongoconn.Database(a.cfg.Mongo.ImageDB),
	)
	if err != nil {
		return nil, 0, err
	}

	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(actID+"_"+imageID, &buf)
	if err != nil {
		return nil, 0, err
	}

	a.log.Debug().Msgf("File size to download: %v\n", dStream)
	return &buf, dStream, nil
}

func (a *ActV1) SoftDeleteActByID(id string) (err error) {
	data, err := a.GetActByID(id)
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

	if a.db == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = a.orm.Update("acts", data)
	if err != nil {
		return err
	}

	return
}

func (a *ActV1) HardDeleteActByID(id string) (err error) {
	conn := *a.db
	if a.db == nil {
		return db.ErrDBConnNotEstablished
	}

	_, err = conn.Exec(conn.Rebind("DELETE FROM production.acts WHERE id=$1"), id)
	if err != nil {
		return err
	}

	return nil
}
