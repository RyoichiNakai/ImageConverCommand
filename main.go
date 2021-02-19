package main

import (
	"conversion_command/check"
	"conversion_command/convator"
	"conversion_command/model"
	"conversion_command/searcher"
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

	extErr := check.Ext(args)
	if extErr != nil {
		fmt.Fprintf(os.Stderr, extErr.Error())
		os.Exit(1)
	}

	fileList, dirErr := searcher.Search(args)
	if dirErr != nil {
		fmt.Fprintf(os.Stderr, dirErr.Error())
		os.Exit(1)
	}

	convator.Convert(args, fileList)

}
