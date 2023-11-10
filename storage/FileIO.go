package storage

import (
	"os"
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
