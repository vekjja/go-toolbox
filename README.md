Go Tool Box
---
A set of golang utility functions

# Compiling
## Using Go Locally
### Prerequisites 
* [golang 1.15 or later](https://golang.org/doc/install)  
  * `brew update && brew install go`  
 
### check_beelzebub_open_files  
`GOOS=linux go build -o ./bin/check_beelzebub_open_files ./check_beelzebub_open_files` 

### desiredApp  
`GOOS=linux go build -o ./bin/desiredApp ./desiredApp` 

## Using Docker
### Prerequisites 
* [docker 18.09 or later](https://docs.docker.com/get-docker/)  
  * `brew update && brew cask install docker` 

### desiredApp  
`DOCKER_BUILDKIT=1 docker build -o bin --build-arg APP=desiredApp .`

#### Specify Target Platform
`DOCKER_BUILDKIT=1 docker build -o bin --build-arg APP=desiredApp --platform linux/386 .`  

These commands will compile the specified application and place the executable in the `./bin` directory. The binary can then be distributed as needed
   

# Tests
## Run test Locally
### Prerequisites 
* [golang 1.15 or later](https://golang.org/doc/install)  
  * `brew update && brew install go` 
```sh
go test -v ./gtb
```
#### Expected Output
```sh 
=== RUN   TestOpenBrowser
/usr/bin/open https://google.com
--- PASS: TestOpenBrowser (0.00s)
=== RUN   TestGetUserInfo
--- PASS: TestGetUserInfo (0.00s)
=== RUN   TestSplitMulti
--- PASS: TestSplitMulti (0.00s)
=== RUN   TestAToUint32
--- PASS: TestAToUint32 (0.00s)
=== RUN   TestGetFilesInDir
--- PASS: TestGetFilesInDir (0.00s)
PASS
ok      github.com/seemywingz/gotoolbox/gtb     0.013s
```
## Run in Docker
### Prerequisites 
* [docker 18.09 or later](https://docs.docker.com/get-docker/)  
  * `brew update && brew cask install docker`
```sh
docker compose up
```
#### Expected Output
```sh 
Starting gotoolbox_go-test_1 ... done
Attaching to gotoolbox_go-test_1
go-test_1  | === RUN   TestOpenBrowser
go-test_1  | xdg-open https://google.com
go-test_1  | --- PASS: TestOpenBrowser (0.00s)
go-test_1  | === RUN   TestGetUserInfo
go-test_1  | --- PASS: TestGetUserInfo (0.00s)
go-test_1  | === RUN   TestSplitMulti
go-test_1  | --- PASS: TestSplitMulti (0.00s)
go-test_1  | === RUN   TestAToUint32
go-test_1  | --- PASS: TestAToUint32 (0.00s)
go-test_1  | === RUN   TestGetFilesInDir
go-test_1  | --- PASS: TestGetFilesInDir (0.01s)
go-test_1  | PASS
go-test_1  | ok         gtb     0.013s
gotoolbox_go-test_1 exited with code 0
```
