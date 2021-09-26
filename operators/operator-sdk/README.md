# Info 
operator-sdk version: "v1.12.0", commit: "d3b2761afdb78f629a7eaf4461b0fb8ae3b02860", kubernetes version: "1.21", go version: "go1.16.7", GOOS: "linux", GOARCH: "amd64"

# Run the example

Guide [https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/](https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/)

### 1- Install operator-sdk 

```bash
export ARCH=$(case $(uname -m) in x86_64) echo -n amd64 ;; aarch64) echo -n arm64 ;; *) echo -n $(uname -m) ;; esac)
export OS=$(uname | awk '{print tolower($0)}')
export OPERATOR_SDK_DL_URL=https://github.com/operator-framework/operator-sdk/releases/download/v1.12.0
curl -LO ${OPERATOR_SDK_DL_URL}/operator-sdk_${OS}_${ARCH}
chmod +x operator-sdk_${OS}_${ARCH} && sudo mv operator-sdk_${OS}_${ARCH} /usr/local/bin/operator-sdk
```

### 2- Basic test

Create project:

```bash
operator-sdk init --domain example.com --repo example.com/hello
```

Create the api:

```bash
operator-sdk create api --group hellogroup --version v1 --kind Hello --resource --controller
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

De Facto standard


## CONS

Same as Kubebuilder
