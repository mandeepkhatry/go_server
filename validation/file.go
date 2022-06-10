package validation

import (
	"io"

	"github.com/go_server/constant"
	"github.com/go_server/logger"
	"github.com/go_server/utils"
)

var ImageTypes = map[string]bool{
	"image/png":  true,
	"image/jpeg": true,
}

type FileValidationParams struct {
	Size   int
	Seeker io.ReadSeeker
}

func FileSize(params FileValidationParams) bool {
	if params.Size > constant.MaxFileUploadSize {
		return false
	}
	return true
}

func MimeType(params FileValidationParams) bool {
	mimeType, _ := utils.GetContentType(params.Seeker)
	logger.Info("", logger.Field("mime", mimeType))
	if _, ok := ImageTypes[mimeType]; !ok {
		return false
	}
	return true
}
