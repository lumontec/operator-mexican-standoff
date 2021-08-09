# Run the example

Guide [https://metacontroller.github.io/metacontroller/guide/create.html](https://metacontroller.github.io/metacontroller/guide/create.html)

### 1- Install metacontroller

```bash
kubectl apply -k https://github.com/metacontroller/metacontroller/manifests/production
```

### 2- Basic test

First of all create the namespace:
```bash
kubectl create namespace hello
```

Define custom resource definition *crd.yaml* and kubectl apply it:

```bash
kubectl apply -n hello -f ./crd.yaml
```
Define the custom controller.yaml and apply it:
```bash
kubectl apply -n hello -f ./controller.yaml
```
Add the webhook implementations as a pure configMap:
```bash
kubectl create -n hello configmap hello-controller --from-file=sync.py
```
Deploy the webhook applying it:
```bash
kubectl apply -n hello -f ./webhook.yaml
```
Now create and update the custom resource to test our controller:

Create
```bash
kubectl -n hello apply -f hello.yaml
kubectl -n hello get pods -a
kubectl -n hello logs your-name
```
Update
```bash
kubectl -n hello patch helloworld your-name --type=merge -p '{"spec":{"who":"My Name"}}'
kubectl -n hello logs your-name
```

## PROS

High level approach, generic controller can be implemented through custom webhooks as http server in any language


## CONS

Support seems quite small
Exploits configmaps for loading the code, not very clean
Requires installation of additional metacontroller controller itself

