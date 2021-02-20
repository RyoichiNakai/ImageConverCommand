package check

import (
	"conversion_command/model"
	"github.com/pkg/errors"
)

var extensions = [6]string{"jpg", "jpeg", "png", "gif", "tiff", "bmp"}

// ここに宣言されると定数になり、再びこのパッケージの関数を呼び出されても値が変わらない

func Ext(args model.Args) error {
	var before bool
	var after bool

	for _, ext := range extensions {
		if *args.BeforeExt == ext {
			*args.BeforeExt = "." + *args.BeforeExt
			before = true
		}

		if *args.AfterExt == ext {
			*args.AfterExt = "." + *args.AfterExt
			after = true
		}
	}

	if !before || !after {
		return errors.New("Extension Error: You must choose image extensions")
	}

	return nil
}
