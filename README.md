# golang - kubernetes - Elastic Cantainer Service 
# ECS 
- ECS is managed container orchestrator
- like docker swarm , kubernetes
- manage life cycle of container create / restart / destroy
- Deploy & load-balance application across multiple servers
- Autoscaling to handle variance in traffic
- Rolling out changes to application

# why ECS
- As compared to ECS, The Kubernetes , Doker Swarm and HashiCrop Nomad etc they all are more effort to get up and running and they are not exactly the simplest or more intuitive to work with especially with kubernetes it is a beast 
- ECS is a simple alternative to these orchestration so that we can basically just use a simple GUI put couples of details that we want our application to operate and AWS is going to handle the rest 

# ECS Infrastructure and Launch 
 ECS has two different launch type 
- EC2 which is a compute service 
- FARGATE which is Serverless way 
 
 Cluster 
- Underlying resources that are ECS intances that can deploy your application it is a physical infractructure

 EC2 Launch 
- We have to manage the underlying EC2 instances 
- We have ECS control plane which is brain behind the operation on container
- We have Cluster
- We have manage dependency on EC2 we have to install Docker to run containers 
- We have to install the ECS agents so we can connect the control plane to instance
- Installing firewall to access certain ports
- Apply routine patches so that our servers all up to date
- The only think ECS does is manage the containers it deploys the containers onto those ec2 instances


