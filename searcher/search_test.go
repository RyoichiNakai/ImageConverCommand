package searcher

import (
	"conversion_command/model"
	"os"
	"reflect"
	"testing"
)

type TestCase struct {
	name string
	in   model.Args
	out  interface{}
}

var (
	cd, _  = os.Getwd()
	before = ".jpg"
	after  = ".jpg"
)

var errCases = []TestCase{
	{
		name: "異常系_引数なし",
		in: model.Args{
			Dir:       "",
			BeforeExt: &before,
			AfterExt:  &after,
		},
		out: "Failed to load directory: no such file or directory",
	},
	{
		name: "異常系_引数あり_存在パス指定",
		in: model.Args{
			Dir:       "notExistDir",
			BeforeExt: &before,
			AfterExt:  &after,
		},
		out: "Failed to load directory: no such file or directory",
	},
}

var normalCases = []TestCase{
	{
		name: "正常系_引数あり_フラグなし",
		in: model.Args{
			Dir:       cd[:len(cd)-len("searcher")],
			BeforeExt: &before,
			AfterExt:  &after,
		},
		out: []string{
			cd[:len(cd)-len("searcher")] + "image/sample1" + before,
			cd[:len(cd)-len("searcher")] + "image/sample2" + before,
			cd[:len(cd)-len("searcher")] + "image/sample3" + before,
		},
		// ファイルパス関連はテストにしやんほうがいいと思う
	},
}

func TestAllCases(t *testing.T) {
	for _, c := range errCases {
		t.Run(c.name, func(t *testing.T) {
			testError(t, c)
		})
	}

	for _, c := range normalCases {
		t.Run(c.name, func(t *testing.T) {
			testNormal(t, c)
		})
	}
}

func testError(t *testing.T, c TestCase) {
	t.Helper() // 原因の特定
	actual, err := Search(c.in)
	if actual != nil {
		t.Errorf("actual expected nil")
	}
	if err.Error() != c.out {
		t.Errorf("case:%v => {Dir: '%s', BeforeExt: '%s', AfterExt:'%s'}, want %q, actual %q", c.name, c.in.Dir, *c.in.BeforeExt, *c.in.AfterExt, c.out, err)
	}
}

func testNormal(t *testing.T, c TestCase) {
	t.Helper() // 原因の特定
	actual, err := Search(c.in)
	if err != nil {
		t.Errorf("err expected nil: %v", err.Error())
	}
	if actual == nil {
		t.Error("Search expected Nonnil")
	}
	//構造体の中身ごと一致するか比較
	if !reflect.DeepEqual(actual, c.out) {
		t.Errorf("case:%v => {Dir: '%s', BeforeExt: '%s', AfterExt:'%s'}, want %q, actual %q", c.name, c.in.Dir, *c.in.BeforeExt, *c.in.AfterExt, c.out, actual)
	}
}
