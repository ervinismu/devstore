package service

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type UploaderService struct {
	CloudName    string
	ApiKey       string
	ApiSecret    string
	UploadFolder string
}

func NewUploaderService(CloudName string, ApiKey string, ApiSecret string, UploadFolder string) *UploaderService {
	return &UploaderService{
		CloudName:    CloudName,
		ApiKey:       ApiKey,
		ApiSecret:    ApiSecret,
		UploadFolder: UploadFolder,
	}
}

func (svc *UploaderService) UploadImage(input *multipart.FileHeader) (string, error) {
	cld, _ := cloudinary.NewFromParams(svc.CloudName, svc.ApiKey, svc.ApiSecret)

	var ctx = context.Background()

	file, _ := input.Open()

	uploaded, err := cld.Upload.Upload(
		ctx,
		file,
		uploader.UploadParams{Folder: svc.UploadFolder},
	)
	if err != nil {
		return "", err
	}

	return uploaded.SecureURL, nil
}
