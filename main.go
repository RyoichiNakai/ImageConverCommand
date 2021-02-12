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
	args.BeforeExt = flag.String("b", "jpg", "Extension before conversion")
	args.AfterExt = flag.String("a", "png", "Extension after conversion")
	flag.Parse()
	args.Dir = flag.Arg(0)

	*args.BeforeExt = "." + *args.BeforeExt
	*args.AfterExt = "." + *args.AfterExt

	if f, err := os.Stat(args.Dir); os.IsNotExist(err) || !f.IsDir() {
		fmt.Fprintf(os.Stderr, "Failed to load directory:\n%v", err)
	}

	fileList := directory_walk.Search(args)
	replace_extension.Convert(args, fileList)
}