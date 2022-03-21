# Get2Json
This `get2json` command-line tool mainly output the header information as Json format when input URLs.

# Features
- `get` command: Be able to output header information by inputting one or multiple URLs.


# How to use

## Run in Docker:
Required
- Install docker

### Run process
Step1: Pull docker image(borischen0203/get2josn)
```bash
docker pull borischen0203/get2josn
```
Step2:  Run docker image as below command
```bash
docker run -it --rm borischen0203/get2josn get
```

### Docker run demo
```bash
# Display the current time in the human text without input parameter
$ docker run -it --rm borischen0203/get2josn get
$ Seven past two

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
```bash
./bin/get2json
```

### Local run demo
```bash
# Display the current time in the human text without input parameter
$ ./bin/get2json get
$

# Display the human text with input numeric time
$ ./bin/get2json get
$
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