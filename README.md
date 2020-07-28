# timer-test
Factorial Program in GO language that can be executed:  
  
(a) on your Mac OS  
(b) on a Docker container on Mac OS  
(c) on Pivotal Cloud Foundry  
(d) on Pivotal Cloud Foundry using a Docker Image  
  
Start with option (a) to make sure you download the program from github  
  
# (a) timer-test on Mac OS  
  
- Prerequisite: You'll need to install the GO Language on your Mac.  
- Do you already have GO installed? Try `Mac $ go version`   
- You should see something like this:   `go version go1.14 darwin/amd64` 
- If you need to install the GO Language, do this: `Mac $ brew install go` 
  
Open a terminal window on your Mac and execute the following command:  

```
Mac $ cd /work  
Mac $ git clone https://github.com/rm511130/timer-test  
Mac $ cd /work/timer-test
Mac $ go run timer-test.go  
```
  
- You should see a message like this one:  `2020/03/14 15:55:55 Starting Timer Test...`
- You can then test it using a browser:    `http://localhost:3000/10`
- And you'll get as a reply:               `I counted to 10,000,000,000 in 2.601754 seconds` 
  
# (b) timer-test using Docker on Mac OS  
  
- Prerequisite: You'll need to install Docker on your Mac.  
- Do you already have Docker installed? Try `Mac $ docker version`  
- You should see both client and server version information, e.g.: `version 19.03.5 for both client and server.`  
- To install Docker on your Mac OS follow the instructions: https://docs.docker.com/engine/installation/mac/  
- To run the Docker Server on your Mac, perform a Mac Spotlight Search for "Docker" and run it.  
  
On an open terminal window with the familiar Docker Whale icon displayed somewhere at the top of the screen of your Mac, execute the following command:  
  
```
Mac $ cd /work
Mac $ git clone https://github.com/rm511130/timer-test
Mac $ cd /work/timer-test
Mac $ docker build -t timer-test .  
Mac $ docker run --publish 6060:3000 --name timer-test --rm timer-test  
```

- You should see a message like this one:  `2020/03/14 15:55:55 Starting Timer Test...`
- You can then test it using a browser:    `http://localhost:6060/10`
- And you'll get as a reply:               `I counted to 10,000,000,000 in 2.601754 seconds` 

  
# (c) timer-test on Pivotal Cloud Foundry  
  
- Prerequisite: You'll Go on your Mac  
Do you already have Godep installed? Try $ godep version  
You should see something like this:        godep v29 (darwin/amd64/go1.6.2)  
To install Godep on your Mac, do this:   $ go get github.com/tools/godep  
  
Open a terminal window on your Mac and execute the following command:  
  
$ cd /work/timer-test  
$ godeps save  
$ ls -a  
.		..		.git		Dockerfile	Godeps		Procfile	README.md	timer-test.go  
$ cf push timer-test -b https://github.com/cloudfoundry/go-buildpack  
  
You should see the usual creating app, route, binding, uploading ... and: urls: timer-test.cfapps.io  
You can now test it:    http://timer-test.cfapps.io/10  


