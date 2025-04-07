# Simple Go Server With DockerFile

This is just a simple go server, for integrating with docker environment.
/
To use this first you need go installed on your system
After that you can clone or download this repository using git clone
```
git clone https://github.com/r4j3sh-com/simple-go-server-with-docker.git
```
Now you can build the binary according to your environment.
I am using here Debian Linux os on dockerfile so the binary need to be build for that environment.
We can do that using:
```
GOOS=linux GOARCH=amd64 go build
```
It will create a binary in current directory named `simple-go-server-with-docker`

Now you can build the docker image and run
To build the docker image:
```
docker build . -t simple-go-server-with-docker:latest
```
To run the docker:
```
docker run -p 8080:8080 simple-go-server-with-docker
```
You will see cli output like `Server started on :8080`

Now opening browser for url <https://localhost:8080> you will see the server is running.


