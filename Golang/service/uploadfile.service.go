package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"strings"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"github.com/google/uuid"
)

const (
	// Create a new Blob Service Client using your Azure Storage account credentials
	accountName   = "feelme"
	accountKey    = "YInTKgO30iWulle6Q5GvUCBJnZG7A+H9MNHp22PmvaWZozjff9J3o86OT01+d9AezbqpIyC8Gw32+AStPonhyg=="
	containerName = "feelme-image/profile"
)

func UploadService(dirName, fileName string, file multipart.File) (string, error) {
	arr := strings.Split(fileName, ".")
	newFileName := uuid.New().String() + "." + arr[len(arr)-1]

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return err.Error(), err
	}

	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	u, _ := url.Parse(
		fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName),
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
