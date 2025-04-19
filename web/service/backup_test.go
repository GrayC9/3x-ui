package service

import (
<<<<<<< HEAD
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
=======
	"os"
	"testing"
)

func TestBackupDB(t *testing.T) {
	testDBPath := "/home/trd12/GolandProjects/3x-ui/db/x-ui.db"
	backupDir := "/home/trd12/GolandProjects/3x-ui/backupplace"

	file, err := os.Create(testDBPath)
	if err != nil {
		t.Fatalf("Ошибка создания тестовой БД: %v", err)
	}
	file.Close()
	//defer os.Remove(testDBPath)

	err = BackupDB(testDBPath, backupDir)
	if err != nil {
		t.Errorf("Ошибка выполнения BackupDB: %v", err)
>>>>>>> 2305d347b6735322e5e3b1f438736b11f74e58ba
	}
}
