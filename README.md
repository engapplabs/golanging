# golanging
## To install on linux
```
+ download go package
+ https://golang.org/dl/
+ then go to Downloads and type on terminal
$ sudo tar -xzvf go1.10.1.linux-amd64.tar.gz -C /usr/local/
$ cd /usr/local/go
$ cd
+ creating workspace
$ mkdir go
$ cd go
$ mkdir src pkg bin
$ cd
+ setting up env vars
+ open up .profile and append at it:
export GOPATH=~/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
+ close and type it
$ source .profile
$ clear
+ to test, type go env
´´´
## First program
´´´
+ go till the folder src we've created inside go
+ create a folder called hello_world
+ create a file called main.go
+ open it up using sublime and type:

package main 

import "fmt"

func main() {
	fmt.Println("CRB greatest of Alagoas!!")
}

+ to junt run, type : go run main.go
+ to build executable: go build main.go, to run it, ./main
´´´

