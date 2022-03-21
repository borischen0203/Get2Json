<img src="https://raw.githubusercontent.com/scraly/gophers/main/gopher-johnny.jpg" alt="gopher-johnny" width=400>

<p align="Left">
  <p align="Left">
    <a href="https://github.com/borischen0203/litclock/actions/workflows/go.yml"><img alt="GitHub release" src="https://github.com/borischen0203/Get2Json/actions/workflows/go.yml/badge.svg?logo=github&style=flat-square"></a>
  </p>
</p>


# Get2Json
This `get2json` command-line tool mainly output the header information as Json format when input URLs.

# Features
- `get` command: Be able to output header information by inputting one or multiple URLs.


# How to use

## Run in Docker:
Required
- Install docker

### Run process
Step1: Pull docker image(borischen0203/get2json)
```bash
docker pull borischen0203/get2json
```
Step2:  Run docker image as below command
```bash
docker run -it --rm borischen0203/get2json get
```

### Docker run demo
```bash
# Display the output header info when input URLs list
$ docker run -it --rm borischen0203/get2json get
$ Please enter URLs:
$ http://www.bbc.co.uk/iplayer
$ http://checkip.amazonaws.com
$ q #<--type[q + enter key] to quit input and run

$ Result:
{
   "Url": "http://www.bbc.co.uk/iplayer",
   "Status-Code": 301,
   "Content-Length": 169
}
{
   "Url": "http://checkip.amazonaws.com",
   "Status-Code": 200,
   "Content-Length": 15
}
```


## Run in Local:

Required
- Install go(version >= 1.7)
- Install `make` cli(https://formulae.brew.sh/formula/make)
```bash
brew install make
```

### Run process
Step1: Clone the repo
```bash
git clone https://github.com/borischen0203/Get2Json.git
```
Step2: Use `make` to execute makefile run test and build
```bash
make all
```
Step3: Execute build file with get command
```bash
./bin/get2json get
```

### Local run demo
```bash
# Display the output header info when input URLs list
$ ./bin/get2json get
$ Please enter URLs:
$ http://www.bbc.co.uk/iplayer
$ http://checkip.amazonaws.com
$ q #<--type[q + enter key] to quit input and run

$ Result:
{
   "Url": "http://www.bbc.co.uk/iplayer",
   "Status-Code": 301,
   "Content-Length": 169
}
{
   "Url": "http://checkip.amazonaws.com",
   "Status-Code": 200,
   "Content-Length": 15
}
```

## Tech stack
- Golang
- Cobra
- Docker
- make
- Github actions
- shell
- go Concurrency(WaitGroup)

## Todo:
- [ ] Be able to input file path and output json file.