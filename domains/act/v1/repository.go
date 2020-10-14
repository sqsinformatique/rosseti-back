package actv1

import (
	"bufio"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"strconv"

	"github.com/sqsinformatique/rosseti-back/internal/db"
	"github.com/sqsinformatique/rosseti-back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func (a *ActV1) actsDB() *mongo.Collection {
	return a.mongodb.Database(a.cfg.Mongo.ActsDB).Collection("acts")
}

func (a *ActV1) CreateAct(request *models.Act) (*models.Act, error) {
	data := *request

	data.CreateTimestamp()

	result, err := a.orm.InsertInto("acts", &data)
	if err != nil {
		return nil, err
	}

	_, err = a.actsDB().InsertOne(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return result.(*models.Act), nil
}

func (a *ActV1) GetActByID(id string) (data *models.Act, err error) {
	actID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}

	data = &models.Act{}

	if a.db == nil {
		return nil, db.ErrDBConnNotEstablished
	}

	err = a.db.Get(data, "select * from production.acts where id=$1", actID)
	if err != nil {
		return nil, err
	}

	a.log.Debug().Msgf("acts %+v", data)

	filter := bson.D{{"id", id}}

	var result models.Act
	err = a.actsDB().FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	data.Body = result.Body

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

func (a *ActV1) CreateImages(actID string, multipartForm *multipart.Form) error {
	for _, fileHeaders := range multipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			bucket, err := gridfs.NewBucket(
				a.mongodb.Database(a.cfg.Mongo.ImageDB),
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
