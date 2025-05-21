# GENV - tiny library for convenient use ENV

### include:

- **_github.com/joho/godotenv_**

### Example:

1. #### Init - Convenient init .env file
   - ```golang
     package main
        
         import (
             "github.com/Alice088/genv"
         )
        
         func main() {
             genv.Init("./env") // if any error == panic(err)
        
             id, err := genv.Get("ID")
             if err != nil {
                 ...
             }
            
             ...
         }

2. #### MuGet - Convenient convert type of ENV value to `int || bool || []byte`
   - ```golang
       package main
    
       import (
           "fmt"
           "github.com/Alice088/genv"
       )
    
       func main() {
           id, err := genv.MuGet[int]("ID") // ID="..." (string)
           if err != nil {
               panic(err)
           }
    
           fmt.Printf("ID: %T", id) //ID: int
       }