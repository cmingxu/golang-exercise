package main

import (
  "log"
  "fmt"
  "strings"

  "github.com/go-ego/gpy"
  "github.com/yanyiwu/gojieba"
)

func main() {
  hans := "中国人香港人，中国人不打香港人"


  log.Println(hans)
  log.Println()

  a := gpy.NewArgs()
  log.Println(gpy.Pinyin(hans, a))

  a.Style = gpy.Tone
  log.Println(gpy.Pinyin(hans, a))

  a.Style = gpy.Tone2
  log.Println(gpy.Pinyin(hans, a))

  var words []string
  //use_hmm := true
  x := gojieba.NewJieba()
  defer x.Free()

  words = x.CutAll(hans)
  fmt.Println(hans)
  fmt.Println("全模式:Println", strings.Join(words, "/"))
  words = x.CutForSearch(hans, true)
  fmt.Println("搜索引擎:Println", strings.Join(words, "/"))


  s := "区块链"
  words = x.Tag(s)
  fmt.Println(s)
  fmt.Println("词性", strings.Join(words, ","))

  s = "棒球"
  words = x.Tag(s)
  fmt.Println(s)
  fmt.Println("词性", strings.Join(words, ","))
}


