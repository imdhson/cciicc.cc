package storage

import (
	"os"
	"path/filepath"
	"strings"
)

func LogOpenFile(address string) (*os.File, error) {
	return os.OpenFile(address, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}
func CreateFile(address string) (*os.File, error) {
	file, f_err := os.Create(address)
	return file, f_err
}

func OpenFile(address string) (*os.File, error) {
	file, f_err := os.Open(address)
	return file, f_err
}

func DeleteAll_space_qr() error {
	remove_err := os.RemoveAll("wwwfiles/assets/space_qr")
	return remove_err
}

func MakeDir_space_qr() error {
	mkdir_err := os.Mkdir("wwwfiles/assets/space_qr", 0755)
	return mkdir_err
}
func Delete_space_qr(space_id string) error {
	//QR코드제거
	file_remove_err := os.Remove("wwwfiles/assets/space_qr/" + space_id + ".png")
	return file_remove_err
}

func Delete_space_host_file(space_id string) error {
	err := filepath.Walk("wwwfiles/host_file/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasPrefix(info.Name(), space_id+".") { // 파일이름이 spaceid. 으로 시작하는것을 발견하면
			file_remove_err := os.Remove(path)
			if file_remove_err != nil {
				return file_remove_err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
