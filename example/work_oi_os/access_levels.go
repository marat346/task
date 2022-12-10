package main

import (
	"os"
  "fmt"
  "bufio"
  )

func main() {
  file,err := os.Create("some.txt")
  if err := os.Chmod("some.txt",0444); err != nil{
    fmt.Println(err)
  }
  writer := bufio.NewWriter(file)
  if err != nil{
    fmt.Println(err)
    return
  }
  defer file.Close()

  writer.WriteString("Say hi")
  writer.WriteString("\n")
  writer.WriteRune('a')
  writer.WriteString("\n")
  writer.WriteByte(67) // C
  writer.WriteString("\n")
  writer.Write([]byte {65,66,67}) //ABC
  writer.WriteString("\n")
  writer.Flush()
}