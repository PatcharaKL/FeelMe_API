package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"os"
	"strings"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func UploadService(dirName, fileName string, file multipart.File) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return err.Error(), err
	}
	accountName := os.Getenv("ACCOUNT_BOLB_NAME")
	accountKey := os.Getenv("ACCOUNT_BOLB_KEY")
	arr := strings.Split(fileName, ".")
	newFileName := uuid.New().String() + "." + arr[len(arr)-1]

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return err.Error(), err
	}

	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	u, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, dirName),
	)
	containerURL := azblob.NewContainerURL(*u, p)

	// Upload file to Azure Blob Storage
	blobURL := containerURL.NewBlockBlobURL(newFileName)
	_, err = azblob.UploadStreamToBlockBlob(context.Background(), file, blobURL, azblob.UploadStreamToBlockBlobOptions{})
	if err != nil {
		return err.Error(), err
	}
	// Get URL of uploaded file
	url := blobURL.URL()
	return url.String(), nil
}
