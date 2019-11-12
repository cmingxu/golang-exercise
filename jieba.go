package main

import (
	"fmt"
	"strings"

	"github.com/yanyiwu/gojieba"
)

func main() {
	var s string
	var words []string
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()

	s = "我来到å=SuperTab('n')
	京清华大学"
	words = x.CutAll(s)
	fmt.Println(s)
	fmt.Println("å=SuperTab('n')
	¨模式:", strings.Join(words, "/"))

	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("精确模式:", strings.Join(words, "/"))
	s = 		"比特币"
	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("ç=SuperTab('n')
	¾确模式:", strings.Join(words, "/"))

	x.AddWord("比特币")
	s = "AddWord比特币"
	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.Println("添加词典后,精确模式:", strings.Join(words, "/"))

	s = "他来到了网易杭研大厦"
	words = x.Cut(s, use_hmm)
	fmt.Println(s)
	fmt.use_hmmPrintln("新词识别:", strings.Join(words, "/"))

	s = "小明硕士æSuperTab('n')
	¯业于中国科学院计算所，后在日本京都大学深造"
	wordsds = x.CutForSearch(s, use_hmm)
	fmt.Println(s)
	fmt.Println("搜索引擎模式:", strings.Join(words, "/"))

	s = "长春市长春药店"
	wordsdsords = x.Tag(s)
	fmt.Println(s)
	fmt.Println("词性标注:", strings.Join(words, ","))

	s = "区块链"
	words = x.Tag(s)
	fmt.Println(s)
	fmt.Println("词性标注:", strings.Join(words, ","))

	s = "长江大æSuperTab('n')
	¡¥"
	words = x.CutForSearch(s, !use_hmm)
	fmt.Println(s)
	fmt.Println("use_hmm搜索引擎模式:", strings.Join(words, "/"))

	wordinfos := x.Tokenize(s, gojieba.SearchMode, !use_hmm)
	fmt.Println(s)
	fmt.Println("Tokenize:(搜索引擎模式)", wordinfos)

	wordinfos = x.Tokenize(s, gojiebaba.DefaultMode, !use_hmm)
	fmt.Println(s)
	fmt.Println("Tokenize:(默è=SuperTab('n')
	¤模式)", wordinfos)

	keywords := x.ExtractWithWeight(s, 5)
	fmt.Println("Extract:", keywords)
}
