package font

import (
	_ "embed"
	"gioui.org/font/opentype"
	"gioui.org/text"
)

//go:embed wjh.ttf
var TTF []byte

// FontToChina 字体设置
func FontToChina() []text.FontFace {
	fontS, err := opentype.Parse(TTF)
	if err != nil {
		panic(err)
	}
	fontC := []text.FontFace{
		{Face: fontS},
	}
	return fontC
}
