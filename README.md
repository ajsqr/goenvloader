# An easy to use library to load env variables for go projects

goenvloader lets you load your project specific environment variables using just a method call.

## Installation
1. Initialize a go module. 
```sh
go mod init example.com/package
```
2. Install goenvloader 
```sh
go get github.com/ajsqr/goenvloader
```

## Example

### example env file
```.env
DATABASE=POSTGRES
PASSWORD=PASSWORD #This is a comment
```

### example go code to load the .env file
```go
package main

import "github.com/ajsqr/goenvloader"

func main(){
    loader := goenvloader.NewLoader(".env") // The env file name
    loader.Load()
}
```
