package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"

    "./models"
)

func main() {
    fmt.Println("Mega Man 1 CRUD Application")
    fmt.Println("1. Create game")
    fmt.Println("2. Read game")
    fmt.Println("3. Update game")
    fmt.Println("4. Delete game")
    fmt.Println("5. Exit")
    var choice int
    fmt.Scan(&choice)
    switch choice {
    case 1:
        createMegaMan1()
    case 2:
        readMegaMan1()
    case 3:
        updateMegaMan1()
    case 4:
        deleteMegaMan1()
    case 5:
        return
    default:
        fmt.Println("Invalid choice")
    }
}

func createMegaMan1() {
    var newGame models.MegaMan1
    fmt.Println("Enter title: ")
    fmt.Scan(&newGame.Title)
    fmt.Println("Enter release date: ")
    fmt.Scan(&newGame.ReleaseDate)
    fmt.Println("Enter developer: ")
    fmt.Scan(&newGame.Developer)
    
    
    b, err := json.Marshal(newGame)
    if err != nil {
        fmt.Println(err)
        return
    }
    
  
    err = ioutil.WriteFile("mega_man_1.json", b, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Game created!")
}

func readMegaMan1() {
  
    b, err := ioutil.ReadFile("mega_man_1.json")
    if err != nil {
        fmt.Println(err)
        return
    }
    

    var game models.MegaMan1
    err = json.Unmarshal(b, &game)
    if err != nil {
        fmt.Println(err)
        return
    }
    

    fmt.Println("Title:", game.Title)
    fmt.Println("Release Date:", game.ReleaseDate)
    fmt.Println("Developer:", game.Developer)
}
func updateMegaMan1() {
  
    b, err := ioutil.ReadFile("mega_man_1.json")
    if err != nil {
        fmt.Println(err)
        return
    }
    

    var game MegaMan1
    err = json.Unmarshal(b, &game)
    if err != nil {
        fmt.Println(err)
        return
    }
    
  
    fmt.Println("Enter new title: ")
    fmt.Scan(&game.Title)
    fmt.Println("Enter new release date: ")
    fmt.Scan(&game.ReleaseDate)
    fmt.Println("Enter new developer: ")
    fmt.Scan(&game.Developer)

    b, err = json.Marshal(game)
    if err != nil {
        fmt.Println(err)
        return
    }
    
err = ioutil.WriteFile("mega_man_1.json", b, 0644)
if err != nil {
fmt.Println(err)
return
}
fmt.Println("Game information updated!")
}

func deleteMegaMan1() {
err := os.Remove("mega_man_1.json")
if err != nil {
fmt.Println(err)
return
}
fmt.Println("Game deleted!")
}







