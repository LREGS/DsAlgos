

Docker containers are a lightweight alternative to virtual machines. Instead of virtualising hardware, docker virtualises the operating system instead and makes the container believe its a seperate operating systme. This alone is often more than enough to help contain back-end web tools.

# Docker Build 

docker build builds a dockerfile. Generally this is kept in a script within the ci/cd pipeline (github actions) so when code is merged the script is fired so it builds the go binary, and then a docker build, and then pushes the docker image onto docker hub.

Pretty cool docker can host your app and then you can access it through the port exposed from the docker image. 

You can stop/kill a contianer. A kill command is a more aggressive way to stop a container

A docker image is the read lonly definition  of a container. A docker container is a virtualised read-write environment. The image describes to docker how to start a new contianer. 

Becuase images are read only they are not actively "running"

some microservices can be defined in an image and can be scaled at a click of a button to spin up multiple containers of hte same image. For example, if you wanted to process 1000 post per second, but your app could only do 100/s. whiwlst you could maybe right a better up, you could also just spin up 10 images that can work in parallel.

# Storage 

many docker containers are stateless. When you create a new continare from the same image it wont store any information from before, when you restart the docker/getting-started container for example you're starting from scratch. :

Although, there is a way to manage persistent state through the use of storage volumes. This is a little bit similar to how a memory card worked in an old playstation. 


docker run -d -e NODE_ENV=development -e url=http://localhost:3001 -p 3001:2368 -v ghost-vol:/var/lib/ghost/content ghost

-e NODE_ENV - tells docker to run this is development mode and not production mode 

url=x 

This tells docker we want it to be accessible at the given url on our local machine 

-p does some port forwarding between the container and our host machine 

the last section mounts the ghost volume which is basically a storage layer for said docker image, as otherwise nothing would be stored

