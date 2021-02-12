package directory_walk

import (
	"conversion_command/model"
	"fmt"
	"os"
	"path/filepath"
)

func Search(args model.Args) (fileList []string) {

	err := filepath.Walk(args.Dir,
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == *args.BeforeExt {
				fileList = append(fileList, path)
			}
			return err
		})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load directory:\n%v", err)
	}

	return
}