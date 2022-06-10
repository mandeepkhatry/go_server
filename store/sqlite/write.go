package sqlite

import (
	"fmt"

	"github.com/go_server/logger"
	"github.com/go_server/models"
	"github.com/google/uuid"
)

func (sq *SQLite) Write(data models.FileData) (string, error) {
	query := fmt.Sprintf(`INSERT INTO data 
			(id,file_name, file_size, content_type) 
			VALUES
			($1,$2,$3,$4);`)

	tx, err := sq.db.Prepare(query)
	if err != nil {
		logger.Debug("", logger.Field("data", data))
		logger.Debug("", logger.Field("error", err))
		return "", err
	}
	var id = uuid.New().String()
	_, err = tx.Exec(id, data.FileName, data.Size, data.ContentType)
	if err != nil {
		logger.Debug("", logger.Field("error", err))
		return "", err
	}
	return id, err
}
