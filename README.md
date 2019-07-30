# AVI Mesh Controller

##  Architecture

The AVI Mesh Controller is a layered collection of independent interoperable units that are used in conjunction to constitute a multi-cluster service mesh using AVI's Enterprise Grade LoadBalancers and Istio.

The controller ingests the Kubernetes API server object updates, Istio's Galley object updates and pumps them into a unified queue. In order to ensure completness of information to formulate an AVI REST API call, every kubernetes or Istio resource is traced back to a Gateway resource. For example, for a kubernetes update of an endpoint, AMC traces to the gateways that are impacted by this change. If the gateway relationship cannot be established, then it's assumed that a gateway resource is not directly or indirectly related to this resource update and that leads to discarding the update for any further action. However, if the gateway can be traced, then a complete walk of the relationships is performed as follows:

    Gateways --> VirtualServices --> Services (Applicable Destination Rules) --> Endpoints --> Servers (based on pod labels)

A bunch of AVI object nodes are generated corresponding to these Istio/K8s objects in the "Object Graph Transformation" layer and these are published to another queue that is eventually consumed by the API execution layer. These AVI object nodes are the `intended` AVI objects that are compared with the AMC cache to arrive at a decision of either CREATE, UPDATE or DELETE a corresponding AVI object via the API execution layer. The API execution layer calls the appropriate REST APIs of the AVI controller to create the corresponding objects in AVI. 

As a consequence to this, the AVI controller  programs the Service Engines deployed in your kubernetes cluster to program all the rules that are expressed in the form of Istio traffic rules. As a first phase of implementation, the AMC works in conjunction with the Istio Gateway functionality by replacing the functionality of Envoy at the edge with AVI's Enterprise grade Service Engines, with a 0 touch object manipulation of the existing Istio Infrastructure. 

The AMC flow can be visualized as follows:

![Alt text](Arch_AVI_Layers.png?raw=true "Title")

## Getting started

##  Build process

The project can be built in a couple of ways:

- [Go Build](#native-go-build)
- [Docker Build](#docker-build)


## Native Go Build

Steps:

    - Configure GOPATH in your machine.
    - mkdir -p $(GOPATH)/src/github/avinetworks/
    - cd $(GOPATH)/src/github/avinetworks/
    - git clone https://github.com/avinetworks/servicemesh
    - cd servicemesh
    - make build

This will generate a binary called: `$(GOPATH)/src/github/avinetworks/sevicemesh/servicemesh`

## Docker build

Steps:

    - Ensure you have docker 17.3 or above that supports docker multi-stage build.
    - git clone https://github.com/avinetworks/servicemesh
    - cd servicemesh
    - make docker

This will generate a docker image by name of `servicemesh:latest`


## Running tests

Steps:

    - git clone https://github.com/avinetworks/servicemesh
    - cd servicemesh
    - make test


## Running the service. 

#### Pre-requisites

The AMC service can be run standalone or inside the kubernetes cluster. If this service is being run outside of the Kubernetes cluster, one would need to expose the Galley server to become accessible over IP. For experiments, this can be achieved by exposing the Galley Service in Istio as NodePort service. Besides this, the following are a list of pre-requisites needed before we get started:

    - A Kubernetes cluster with Istio deployed with MCP services enabled.
    - A AVI controller accesible over IP: 
         -  The AVI controller should be configured with the Kubernetes/Openshift cloud. 
         -  The IPAMs for the North/South and East/West services should be configured.
         -  The service syncing for Backend/Frontend services should be disabled on the cloud.
 
 #### Running AMC outside the cluster

 Assuming you have built the project successfully by following the Build Process section. The following environment variables should be exposed if you are running this service outside the kubernetes cluster:
 
 `export ISTIO_ENABLED=True` - Enables the Istio MCP Client
 
 `export CTRL_USERNAME=<username>` - The AVI Controller username.
 
 `export CTRL_PASSWORD=<password>` - The AVI Controller password
 
 `export MCP_URL=<Galley_URL>:<GalleyNodePort>` - The endpoint to contact Galley.

 `export CTRL_IPADDRESS=<AVI_API_SERVER_ADDR>` - The AVI controller API endpoint with port if applicable.
 
 Post these steps - one can simply start the AMC service using: `./servicmesh`

#### Running AMC inside the cluster

 If you are running inside a kubernetes cluster, then a lot of automation is provided out of the box. Please follow the below steps to run it:
 
 - Clone the code.
 - `cd servicemesh/k8s_tmpl`
 - Edit the secret.yaml file and update it with the relevant information by encoding it to base64. 
 - `kubectl create -f secret.yaml`
 - `kubectl create -f deployment.yaml`
 
 The above should bring up a POD in your kubernetes cluster running the AVI Mesh Controller.
 
