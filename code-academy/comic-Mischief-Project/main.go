package main
import "fmt"

func main() {
  var publisher, writer, artist, title string
  var year, pageNumber int

  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14

  /*
  We’ll need to define our condition grade. This is a score from 0.0 to 10.0 with decimal values. Because of the decimal values, we’ll need to use a float to store it.
  */
  var grade float32
  grade = 6.5

  fmt.Println(
    title, 
    "written by", writer, 
    "drawn by", artist, 
    "Published by", publisher, 
    year, pageNumber, "pages", grade,
  )

  fmt.Println("####################")
  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  publisher = "DizzyBooks Publishing Inc."
  year = 2013
  pageNumber = 160
  grade = 9.0




  fmt.Println(
    title, 
    "written by", writer, 
    "drawn by", artist, 
    "Published by", publisher, 
    year, pageNumber, "pages", grade,
  )
}