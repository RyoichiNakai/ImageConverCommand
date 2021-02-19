package check

import (
	"conversion_command/model"
	"github.com/pkg/errors"
)

var extensions = [6]string{"jpg", "jpeg", "png", "gif", "tiff", "bmp"}
var (
	before = false
	after  = false
)

func Ext(args model.Args) error {
	dot := "."

	for _, ext := range extensions {
		if *args.BeforeExt == ext {
			*args.BeforeExt = dot + *args.BeforeExt
			before = true
		}

		if *args.AfterExt == ext {
			*args.BeforeExt = dot + *args.AfterExt
			after = true
		}
	}

	if !before || !after {
		err := errors.New("Extension Error: You must choose image extensions.")
		return err
	}

	return nil
}
