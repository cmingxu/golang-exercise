package main

import (
  "log"
  "io/ioutil"

  "github.com/yanyiwu/gojieba"
)

func main() {
  textfile := "./x"
  text, _ := ioutil.ReadFile(textfile)
  log.Println(string(text))

  x := gojieba.NewJieba()
  defer x.Free()

  var words []string
  words = x.CutForSearch(string(text), true)
  log.Println(words)
  log.Println(len(words))

  freq := make(map[string]int)

  for _, w := range words {
    if _, ok := freq[w]; ok {
      freq[w] += 1
    }else{
      freq[w] = 0
    }
  }


  for w, c := range freq{
    log.Println(w, "=> ", c)
  }
}

