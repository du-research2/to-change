package tochange

import "github.com/ventu-io/go-shortid"
import "fmt"


func Shalom(name string) string {
    // Return a greeting that embeds the name in a message.
    fmt.Println(shortid.Generate())
    message := fmt.Sprintf("Shalom Shalom, %v. Welcome!", name)
    return message
}


//func main() {
    // Get a greeting message and print it.
//    fmt.Println("Your short ID is:")
//    fmt.Println(shortid.Generate())
//}

