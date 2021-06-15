# Question 1
Create a Hello World executable application, choose the code on any of the following scripting languages, include instructions to run the application.

## Golang
### 1.
```bash
# run
go run hello.go

# create a binary 
go build -o hello hello.go
chmod +x hello 
./hello
```
### 1.1
```bash
# run
go run hello.go Rafael
# build
go build -o hello hello.go
chmod +x hello 
./hello Rafael
```

## Python
### 1.
```
python hello.py
```

### 1.1
```
python hello.py Rafael
```
## Bash
### 1.
```
chmod +x hello.sh
./hello.sh
```

### 1.1
```
chmod +x hello.sh
./hello.sh Rafael
```

# 1.2
Method 1 - Download zip file of the project
```
wget https://github.com/turnes/boldlink-phaseone/archive/refs/heads/main.zip -o rafaelturnes-phaseone.zip
```

Method 2 - Clone repository
```
 git clone git@github.com:turnes/boldlink-phaseone.git
 git archive -o rafaelturnes-phaseone.zip HEAD
```