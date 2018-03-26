
# yogo-rest

The programming language that has been chosen is golang.

## Install and Setup GO
Download and configure your workspace with latest version of Go and correct environment path.

### Last Go version
https://golang.org/dl/

### Windows
http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/

### Linux
http://www.tecmint.com/install-go-in-linux/

## Get Source Code
On windows:
```
cd %GOPATH%/src/github.com/giansalex
```

On Linux:
```
cd $GOPATH/src/github.com/giansalex
```

Create the folders "github.com" and "giansalex" if not already created.

Then:
```
git clone https://github.com/giansalex/yogo-rest
cd yogo-rest
```

## Run the API
```
dep ensure
go run main.go
```
