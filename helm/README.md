# Run the example

Guide [https://helm.sh/docs/intro/install/](https://helm.sh/docs/intro/install/)

### 1- Install helm 
...

### 2- Basic test

Configure a simple helm chart and install it (to test it add --dry-run)
```bash
helm install hello-helm ./hello-helm
```

Get all current charts and read the current chart values
```bash
helm list
...
helm get values hello-helm --all
```

Change value inside ./hello-helm/values.yaml
and upgrade the chart with the new values:
```bash
helm upgrade hello-helm ./hello-helm
```
This will upgrade revision number along with our deployment with the new value


## Pros

High level approach, no code required, basically a configurator for our manifests

## Cons

Cannot upgrade crds
Cannot implement complex flows, such as generating some certificate before doing other steps. Doing this requires usage of go functions inside template files, badly mixing declarative approach with imperative approach (see [https://banzaicloud.com/blog/creating-helm-charts/](https://banzaicloud.com/blog/creating-helm-charts/) for an example)
