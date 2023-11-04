package storage

import "os"

func LogOpenFile(address string) (*os.File, error) {
	return os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}
