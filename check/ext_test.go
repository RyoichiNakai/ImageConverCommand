package check

import (
	"conversion_command/model"
	"testing"
)

type TestCase struct {
	name string
	in   model.Args
	out  interface{}
}

var (
	errFrom    = "jpg"
	errTo      = "png"
	normalFrom = "jpg"
	normalTo   = "png"
	empty      = ""
	errExt     = "aaa"
)

var errCases = []TestCase{
	{
		name: "異常系_フラグ引数1つあり_b='aaa'_a='png'",
		in: model.Args{
			Dir:       empty,
			BeforeExt: &errExt,
			AfterExt:  &errTo,
		},
		out: "Extension Error: You must choose image extensions",
	},
	{
		name: "異常系_フラグ引数1つあり_b='jpg'_a='aaa'",
		in: model.Args{
			Dir:       empty,
			BeforeExt: &errFrom,
			AfterExt:  &errExt,
		},
		out: "Extension Error: You must choose image extensions",
	},
	{
		name: "異常系_フラグ引数2つあり_b='aaa'_a='aaa'",
		in: model.Args{
			Dir:       empty,
			BeforeExt: &errExt,
			AfterExt:  &errExt,
		},
		out: "Extension Error: You must choose image extensions",
	},
}

var normalCases = []TestCase{
	{
		name: "正常系_フラグ引数2つあり_b='jpg'_a='png'",
		in: model.Args{
			Dir:       empty,
			BeforeExt: &normalFrom,
			AfterExt:  &normalTo,
		},
		out: nil,
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
	err := Ext(c.in)
	if err.Error() != c.out {
		t.Errorf("case:%v => {Dir: %q, BeforeExt: %q, AfterExt: %q}, want %q, actual %q", c.name, c.in.Dir, *c.in.BeforeExt, *c.in.AfterExt, c.out, err)
	}
}

func testNormal(t *testing.T, c TestCase) {
	t.Helper() // 原因の特定
	err := Ext(c.in)
	if err != nil {
		t.Errorf("case:%v => {Dir: %q, BeforeExt: %q, AfterExt: %q}, want %q, actual %q", c.name, c.in.Dir, *c.in.BeforeExt, *c.in.AfterExt, c.out, err)
	}
}
