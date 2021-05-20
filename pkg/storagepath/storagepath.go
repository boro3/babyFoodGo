package storagepath

import (
	"os"
	"path/filepath"
)

//Creates file for given path including parent folders if the directory does not exsist
//Path including the file name as input string is required
func Create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}
