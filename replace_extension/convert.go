package replace_extension

import (
	"conversion_command/model"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func Convert(args model.Args, fileList []string) {
	for _, path := range fileList {
		str := encodeImageFile(path)
		err := os.Remove(path) // エンコードしたファイルを削除
		if err != nil {
			fmt.Fprintf(os.Stderr, "Remove File Error: %v\n", err)
		}
		decodeImageFile(args, path, str)
	}
}

func encodeImageFile(path string) string {
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read File Error: %v\n", err)
	}
	defer readFile.Close()

	fi, _ := readFile.Stat() //FileInfo interface
	size := fi.Size()        //ファイルサイズ

	data := make([]byte, size)
	_, err = readFile.Read(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read File Error: %v\n", err)
	}
	fmt.Println(filepath.Base(path) + "を読み込んでいます...")
	return base64.StdEncoding.EncodeToString(data)
}

func decodeImageFile(args model.Args, path, str string) {
	data, _ := base64.StdEncoding.DecodeString(str)
	ext := filepath.Ext(path)
	dstPath := path[:len(path) - len(ext)] + *args.AfterExt

	writeFile, err := os.Create(dstPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Create File Error: %v\n", err)
	}
	defer writeFile.Close()

	n, err := writeFile.Write(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Write File Error: %v\n", err)
	}
	fmt.Println(filepath.Base(dstPath) + "を出力しています...")
	fmt.Println("出力バイト数: " + strconv.Itoa(n) + "byte")
}

