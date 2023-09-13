package common

import (
	"io"
)

type Storage interface {
	ID() uint64

	Exists(path string) (bool, error)
	IsFile(path string) (bool, error)
	IsDir(path string) (bool, error)

	List(dir, marker string, limit uint) ([]string, string, error) // objects, next marker, error
	Delete([]string) error
	Copy(string, string) error
	Move(string, string) error // oss 先 copy 后删除

	PutFile(string, string) (string, error) //  dst, file
	Put(string, io.Reader) (string, error)  // dst, reader
	Append(string, io.Reader, int64) (string, error)

	Get(string) (io.ReadCloser, error)
	GetURL(string) (string, error)
}
