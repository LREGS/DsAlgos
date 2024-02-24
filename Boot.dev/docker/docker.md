

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

A volume is the only way for data to persists between container restarts. If you want to delete your volume you can:

docker run -d -e NODE_ENV=development -e url=http://localhost:3001 -p 3001:2368 -v ghost-vol:/var/lib/ghost/content ghost

docker -rm CONTAINER_ID deletes a container 


# Load Balancers 

There is no way that a single server ( a single computer) could handle all of the requests that come from a website the size of google. 

As a result they use  something called a load balance that can route requests to different servers. 

A central server called a load balance recieves traffic from users then routes those raw requests to different back end application servers. In the case of google this splits the worlds traddic across potentally many differnent thousands of computers. 

One method of load balancing is to make the cuyrrent server more powerful (verticle scaling)

Horizontal scaling is buying more servers 

Scaling horizontaly is way more complex 

With a loadbalancer, it does no real work, all it does is route/point work. It doesnt knowl how to handle requests specifically, but its knows what applications are responsible for the requests. 

So a loadbalance sits on top of multiple back end servers that can execute on the requests. 

A round robin is on e way of doing load balancing. My initial issue with this is that its not filtering requests based on size. so worst case scenario is that one server recieves all the most intense requests, and one server recieves all the least intensive requests 
- for a normal website they will normally be about the same size so it will be fine 

In the event though that request resources are needing to be considered 

A more sophisticated approach involces the servers sending the loadbalancers its current cpu usage. The load balancer can then identify the servers that are using the least amount of cpu and then sent requests to them and keep it more equal. 

Once all your servers get close to their compacity, a new server can be spawned in, and the load balance will prioritise sending load to thew newest servers until it is balanced out again
