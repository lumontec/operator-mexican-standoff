# Operator mexican standoff

Bloody shootout between kubernetes operators to test them on very basic examples.

## Classification

First of all, a simple classification. The project is divided in two folders:
* operators
* package-managers

Given the always misguised topic of helm-vs-operator I meant to divide technologies which clearly aim to solve at different functions (see [https://cloudblogs.microsoft.com/opensource/2020/04/02/when-to-use-helm-operators-kubernetes-ops/](https://cloudblogs.microsoft.com/opensource/2020/04/02/when-to-use-helm-operators-kubernetes-ops/))

## Set up test environment

Create simple local test environment

```bash
kind create cluster --config=./kind-config-3nodes.yaml --image kindest/node:v1.20.0
```


## Test case

Create and update a simple custom resource leveraging the operators

## Custom resource definition

```yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: helloworlds.example.com
spec:
  group: example.com
  version: v1
  names:
    kind: HelloWorld
    plural: helloworlds
    singular: helloworld
  scope: Namespaced
  subresources:
    status: {}
```

HelloWorld is a custom resource that creates a busybox pod echoing the following:
"Hello, <name> !"

## Custom resource

The follwing resource can be kubectl applied to create and edit the resource behavior:

```yaml
apiVersion: example.com/v1
kind: HelloWorld
metadata:
  name: your-name
spec:
  who: Your Name
```

This will create and update pod with name "your-name", that will output to logs:
```text
Hello, Your Name !
```
The pod spec will look like this:
```yaml
---
apiVersion: v1
kind: Pod
metadata:
  name: <pod name selected from the user> 
spec:
  restartPolicy: OnFailure
  containers:
  - name: hello
    image: busybox
    command: ["echo", "Hello, %s!" <text name selected from the user>] 
```


## Simply generate a test cluster

```bash
kind create cluster --config=kind-config-3nodes.yaml --image=kindest/node:v1.20.2
```
