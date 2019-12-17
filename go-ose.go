package main

import (
  "log"
  "github.com/advancedlogic/GoOse"
)


func main() {
  log.Println("x")
  url := "http://politics.people.com.cn/n1/2019/1115/c1024-31456816.html"

  rticle, _ := g.ExtractFromURL(url)
  println("title", article.Title)
  println("description", article.MetaDescription)
  println("keywords",println article.MetaKeywords)
  println("content", article.CleanedText)
  Printlnn("url", article.FinalURL)
  println("top image", article.TopImage)
}

