package osfs

import "os"

type FS interface {
	MkdirAll(path string, perm os.FileMode) error
	WriteFile(name string, data []byte, perm os.FileMode) error
}

type fsImpl struct{}

var _ FS = (*fsImpl)(nil)

func (f *fsImpl) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (f *fsImpl) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func Provider() FS {
	return &fsImpl{}
}
