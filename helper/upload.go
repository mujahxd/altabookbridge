package helper

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/mujahxd/altabookbridge/config"
)

// var URLCloudService = os.Getenv("URLCLOUDRINARY")

func Upload(image *multipart.FileHeader) (string, error) {
	config := config.InitConfig()

	file, err := image.Open()
	if err != nil {
		log.Println("error from reading file appload", err.Error())
		return "", err
	}

	cldService, err := cloudinary.NewFromURL(config.URLCLOUDINARY)
	if err != nil {
		log.Println("error on connection to cldService", err.Error())
		return "", err
	}

	ctx := context.Background()
	responseCloud, err := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	if err != nil {
		log.Println("error from taking response from cloudinary", err.Error())
		return "", err
	}

	book_image := responseCloud.URL

	file.Close()
	return book_image, nil
}
