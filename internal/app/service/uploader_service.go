package service

import (
	"context"
	"mime/multipart"
	"time"

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

// func (svc *UploaderService) UploadImage(input interface{}) (string, error) {
func (svc *UploaderService) UploadImage(imageName string, input *multipart.FileHeader) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(svc.CloudName, svc.ApiKey, svc.ApiSecret)
	if err != nil {
		return "", err
	}

	file, _ := input.Open()
	defer file.Close()

	//upload file
	uploaded, err := cld.Upload.Upload(ctx,
		file,
		uploader.UploadParams{Folder: svc.UploadFolder, PublicID: imageName},
	)
	if err != nil {
		return "", err
	}

	return uploaded.SecureURL, nil
}
