package searcher

import (
	"conversion_command/model"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func Search(args model.Args) ([]string, error) {
	var fileList []string
	err := filepath.Walk(args.Dir,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == *args.BeforeExt {
				fileList = append(fileList, path)
			}
			return err
		})

	if err != nil {
		return nil, errors.New("Failed to load directory: no such file or directory")
	}

	return fileList, nil
}
