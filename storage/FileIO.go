package storage

import "os"

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
