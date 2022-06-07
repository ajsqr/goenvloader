# An easy to use library to load env variables for go projects

goenvloader lets you load your project specific environment variables using just a method call.


## Example

### example env file
```.env
DATABASE=POSTGRES
PASSWORD=PASSWORD #This is a comment
```

### example go code to load the .env file
```go
package main

import "github.com/ajay-ajith/goenvloader"

func main(){
    loader := goenvloader.NewLoader(".env") // The env file name
    loader.Load()
}
```