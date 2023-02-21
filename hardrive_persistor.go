package persistor

import (
	"io/fs"
	"os"
)

// NewHarDrivePersistor creates new HarDrivePersistor.
func NewHarDrivePersistor() HarDrivePersistor {
	return HarDrivePersistor{
		Perm: 0755,
	}
}

// HarDrivePersistor provides for saving named data to the hard drive.
type HarDrivePersistor struct {
	Perm fs.FileMode
}

// Persist saves data to the "path/name" file. If a path is empty saves the
// file to the current directory.
func (persistor HarDrivePersistor) Persist(name string, data []byte,
	path string) (err error) {
	if name == "" {
		return ErrUndefinedName
	}
	if path == "" {
		path = "."
	}
	return os.WriteFile(path+string(os.PathSeparator)+name, data,
		persistor.Perm)
}
