# kubeq
Kubernetes Job Scheduler

TL;DR: Make a job queue system similiar to Resque, that leverages the Kubernetes Job resource.

# Get Started

* Install minikube (https://github.com/kubernetes/minikube/releases)
* Start minikube: `minikube start`
* Fire up redis: `kubectl create -f deployments/redis-pod.yml`
* Create the redis service: `kubectl create -f deployments/redis-service.yml`
* Start the scheduler: `kubectl create -f deployments/kubeq-deploy.yml`
