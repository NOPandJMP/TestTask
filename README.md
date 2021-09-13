# TestTask
This is a test assignment, which I completed for the company IT-Expertiza. My task was as follows : 
To write a simple REST application to keep track of jobs.
Supported at least the following operations: adding, removing and updating a job.
A workstation contains the following information:
   - computer name
   - network address
   - Current user name
The client part runs on windows. Running the client without parameters automatically adds/updates information about the current system. Group operations are welcome.
The server side stores data in JSON format in a PostgreSQL database and runs on linux.
bonus1:
  The client knows how to work on linux and windows systems.
bonus2:
  The server part comes in a deb package and is registered as a systemd service.


I created a client application and small tests for it , these are client.go , client_test.go header.go files . 
Since it is very convenient to transfer data in a query using json, I created a structure UserInfo. In go it is very convenient to work with json , so I use json.Marshal
```Golang
type UserInfo struct {
	PcName      string
	UserName    string
	NetworkAddr string
}
```
To start the client you need to enter in the terminal
```bash
go run client.go header.go new
```
To build and start client
```bash
go build -o client client.go header.go

/client new
```
If we want to build our client for windows, because it works the same way on windows, we need to write :
```bash
GOOS=windows go build -o win_client client.go header.go
```

The server handles post, put, delete requests.
A post request means we need to create and add a new user to the database, put to update the information, delete to delete the user's information. 
The user information is stored in the postgresql database, the job was to save the information in json format, but in my opinion it is much more convenient to work with normal data, moreover we can always make a json.Marshal and send a json response. 
In order for our server to start , we need to install the postgressql server and create a database and a table in it ,
with the following columns pc_name , username , network_addr , it looks like this .
```sql
create table Info(
id session PRIMARY KEY NOT NULL , 
pc_name varchar(64) NOT NULL , 
username varchar(64) NOT NULL , 
network_addr varchar(64) NOT NULL
);
```

When we have everything ready, we can start the server, or build it and run it, this is done as follows. 
```bash
go run server.go header.go handlers.go 

go build -o server server.go header.go handlers.go
```
After that our server will start at: localhost:8000

And we will refer to our api at: http://localhost:8000/api/v1/workspace 

If the request is correct, our server will process and return a successful response, if something goes wrong, the server will return an error and the code 424 . 

# DEB PACKAGE CREATION

First, we create the server folder , then navigate to it 
```bash
mkdir server
cd server/
```
Then move our source files there and compile them 
```bash
go build -o systemd  server.go header.go handlers.go
ls 
go.mod  go.sum  handlers.go  header.go  server.go  systemd
```
We built our server and with the command ls brought out all the files that are in our folder

# MANIFESTO CREATION


We will use the package folder to build the package, so that the program files are not mixed up with the source files and they are not included in the package.

```bash
mkdir -p package/DEBIAN
```

The first thing to look at is the size of the program files, since in this case there is only one file, just look at its size
```bash
du -k ./systemd 
7012	./systemd
```
 Next, you need to figure out which packages your program will depend on. To do this, use the command objdump
 ```bash
 objdump -p ./systemd | grep NEEDED
  NEEDED               libpthread.so.0
  NEEDED               libc.so.6
 ```
 In this case the program needs libc and libpthread. To see which package it is in, run
 ```bash
 dpkg -S libc.so.6
libc6:amd64: /lib/x86_64-linux-gnu/libc.so.6

dpkg -S libpthread.so.0
libc6:amd64: /lib/x86_64-linux-gnu/libpthread.so.0
```

The package is called libc6. Then create a manifest file with the following contents
```bash
vi package/DEBIAN/control
```
Package: systemd
Version: 1.0
Section: unknown
Priority: optional
Depends: libc6
Architecture: amd64
Essential: no
Installed-Size: 7012
Maintainer: github.com/NOPandJMP <email@mail.ru>
Description: REST applications 

# FILE LOCATION

The manifest is ready. Now we need to create a folder structure in the package folder, analogous to what we have in the root file system. In this case, create a folder usr/bin and put the executable there
```bash
mkdir -p package/usr/bin
mv ./systemd package/usr/bin
```

# INSTALLATION SCRIPTS
 
 ```bash
 vi package/DEBIAN/postinst
 
 #!/bin/bash
echo "Hello from github.com/NOPandJMP installed"
```

# PACKAGE ASSEMBLY AND VERIFICATION
 
 It remains to build the configured package. To do this, use the following command

```bash
dpkg-deb --build ./package
```
After completing the build, you can install it with apt
```bash
sudo apt install ~/systemd.deb
```


I would very much like feedback from knowledgeable people to point out my mistakes, and give me advice on how to be better. Thank you all for reading! :) 
