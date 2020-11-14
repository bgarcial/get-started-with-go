package main

import "fmt"

func main() {
  var stationName string
  var nextTrainTime int8
  var isUptownTrain bool
  
  stationName = "Union Square"
  nextTrainTime = 12
  isUptownTrain = false
  
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)
  
  stationName = "Grand Central"
  nextTrainTime = 3
  isUptownTrain = true
  
  fmt.Println("")
  fmt.Println("Current station:", stationName)
  fmt.Println("Next train:", nextTrainTime, "minutes")
  fmt.Println("Is uptown:", isUptownTrain)

  /* Another way to assing a variable  */
  var kilometersToMars int32 = 62100000

  /* Strings */
  nameOfSong = "Stop Stop"
  nameOfArtist = "The Julie Ruin"
  var description string
  description = nameOfSong + " is by: " + nameOfArtist + "."
  fmt.Println(description)	

  var favoriteSnack string
  favoriteSnack = "Wageningen cheese"
  fmt.Println("My favorite snack is " + favoriteSnack + ".")

  var emptyInt int8
  
  var emptyFloat float32
  
  var emptyString string
  
  // Finally, print them all out
  fmt.Println(emptyInt, emptyFloat, emptyString)


}