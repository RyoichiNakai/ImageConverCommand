package main

import (
	"conversion_command/directory_walk"
	"conversion_command/model"
	"conversion_command/replace_extension"
	"flag"
	"fmt"
	"os"
)

func main() {
	var args model.Args
	extensions := [6]string{"jpg", "jpeg", "png", "gif", "tiff", "bmp"}
	args.BeforeExt = flag.String("b", "jpg", "Extension before conversion")
	args.AfterExt = flag.String("a", "png", "Extension after conversion")
	flag.Parse()
	args.Dir = flag.Arg(0)

	var (
		before = false
		after  = false
	)

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
		fmt.Println("Extension Error:")
		fmt.Println("You must choose image extensions.")
		os.Exit(0)
	}

	if f, err := os.Stat(args.Dir); os.IsNotExist(err) || !f.IsDir() {
		fmt.Fprintf(os.Stderr, "Failed to load directory:\n%v", err)
	}

	fileList := directory_walk.Search(args)
	replace_extension.Convert(args, fileList)
	fmt.Println("aaaa")
}
