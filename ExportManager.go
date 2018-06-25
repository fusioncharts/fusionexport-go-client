package FusionExport

import (
	"fmt"
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
	"context" // needed for s3
	"bytes" // needed for s3 and ftp
	"net/http" // needed for s3
	"github.com/aws/aws-sdk-go/aws"  // needed for s3
	"github.com/aws/aws-sdk-go/aws/credentials"  // needed for s3
	"github.com/aws/aws-sdk-go/service/s3" // needed for s3
	"github.com/aws/aws-sdk-go/aws/session" // needed for s3
	"github.com/aws/aws-sdk-go/service/s3/s3manager" // needed for s3
	
)

type ExportManager struct {
	Host string
	Port int
}

func NewExportManager() ExportManager {
	em := ExportManager{
		Host: DEFAULT_HOST,
		Port: DEFAULT_PORT,
	}
	return em
}

func (em *ExportManager) SetConnectionConfig(host string, port int) {
	em.Host = host
	em.Port = port
}

func (em *ExportManager) Export(exportConfig ExportConfig, exportDone func([]OutFileBag, error), exportStateChanged func(ExportEvent)) (Exporter, error) {
	exp := Exporter{
		ExportConfig:              exportConfig,
		ExportDoneListener:        exportDone,
		ExportStateChangeListener: exportStateChanged,
		ExportServerHost:          em.Host,
		ExportServerPort:          em.Port,
	}

	err := exp.Start()

	return exp, err
}

func SaveExportedFiles(fileBag []OutFileBag) error {
	for _, file := range fileBag {
		fileData, err := base64.StdEncoding.DecodeString(file.FileContent)
		if err != nil {
			return err
		}

		dir := filepath.Dir(file.RealName)
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(file.RealName, fileData, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
func GetRegion(bucket string) string {
	sess := session.Must(session.NewSession())
	region, err := s3manager.GetBucketRegion(context.Background(), sess, bucket, "us-west-2")
	if err != nil {
	   return err.Error();
	}
	return region
}
func UploadFileToAmazonS3(fileBag []OutFileBag, bucket string, accessKey string, secretAccessKey string){
	aws_access_key_id := accessKey
	aws_secret_access_key := secretAccessKey
	token := "" 
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token) 
	_, err := creds.Get() 
	if err != nil { 
		fmt.Printf("invalid credentials: %s", err) 
	} 
	region := GetRegion(bucket)
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds) 
	svc := s3.New(session.New(), cfg) 
	for _, file := range fileBag {
		fileData, err := base64.StdEncoding.DecodeString(file.FileContent)
		if err != nil {
			fmt.Printf("%s",err)
			
		}
		size := int64(len(fileData))

		fileBytes := bytes.NewReader(fileData) 
		fileType := http.DetectContentType(fileData) 
		params := &s3.PutObjectInput{ 
		Bucket: aws.String(bucket), 
		Key: aws.String(file.RealName), 
		Body: fileBytes, 
		ContentLength: aws.Int64(size), 
		ContentType: aws.String(fileType), 
		} 
		resp, err := svc.PutObject(params) 
		if err != nil { 
			fmt.Printf("bad response;response: %s, error: %s", resp,err) 
		} 
		fmt.Printf("file uploaded")

		
	} 
}


func GetExportedFileNames(fileBag []OutFileBag) []string {
	var fns []string

	for _, file := range fileBag {
		fns = append(fns, file.RealName)
	}

	return fns
}
