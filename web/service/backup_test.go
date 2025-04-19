package service

import (
	"testing"
)

func TestBackupDBToS3(t *testing.T) {
	testDBPath := "/Users/danilamanakov/GolandProjects/3x-ui/db/x-ui.db"
	awsAccessKey := "admin"
	awsSecretKey := "admin123"
	awsRegion := "fsn1"
	s3Endpoint := "http://localhost:9000"
	bucketName := "my-backup-bucket"

	err := BackupDBToS3(testDBPath, bucketName, s3Endpoint, awsRegion, awsAccessKey, awsSecretKey)
	if err != nil {
		t.Errorf("Ошибка выполнения BackupDBToS3: %v", err)
	}
}
