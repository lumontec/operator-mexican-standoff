# Info 
Version: main.version{KubeBuilderVersion:"3.1.0", KubernetesVendor:"1.19.2", GitCommit:"92e0349ca7334a0a8e5e499da4fb077eb524e94a", BuildDate:"2021-05-27T17:54:28Z", GoOs:"linux", GoArch:"amd64"}

# Run the example

Guide [https://book-v1.book.kubebuilder.io/](https://book-v1.book.kubebuilder.io/)

### 1- Install kubebuilder cli

```bash
curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH)
chmod +x kubebuilder && mv kubebuilder /usr/local/bin/
```

### 2- Basic test

Create project:

```bash
kubebuilder init --domain example.com --repo example.com/hello
```

Create the api:

```bash
kubebuilder create api --group hellogroup --version v1 --kind Hello
```

* Configure our crd spec insde ./hello/api/v1/hello_types.go
* Implement controller inside ./hello/controllers/hello_controller.go


Install all the resources in the cluster:
```bash
make install
```

run the example
```bash
make run
```




## PROS

High level approach, generic controller can be implemented through custom webhooks as http server in any language


## CONS

Support seems quite small
Exploits configmaps for loading the code, not very clean
Requires installation of additional metacontroller controller itself

