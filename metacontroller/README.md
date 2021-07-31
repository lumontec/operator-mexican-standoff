# Run the example

Guide [https://metacontroller.github.io/metacontroller/guide/create.html](https://metacontroller.github.io/metacontroller/guide/create.html)

### 1- Install metacontroller

```bash
kubectl apply -k https://github.com/metacontroller/metacontroller/manifests/production
```

### 2- Basic test

Create namespace hello:
```bash
kubectl create namespace hello
```
Define custom resource definition *crd.yaml* and kubectl apply it:
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
Define the custom controller.yaml and apply it:
```yaml
apiVersion: metacontroller.k8s.io/v1alpha1
kind: CompositeController
metadata:
  name: hello-controller
spec:
  generateSelector: true
  parentResource:
    apiVersion: example.com/v1
    resource: helloworlds
  childResources:
  - apiVersion: v1
    resource: pods
    updateStrategy:
      method: Recreate
  hooks:
    sync:
      webhook:
        url: http://hello-controller.hello/sync
```
Implement the webhook in sync.py:
```python
from http.server import BaseHTTPRequestHandler, HTTPServer
import json

class Controller(BaseHTTPRequestHandler):
  def sync(self, parent, children):
    # Compute status based on observed state.
    desired_status = {
      "pods": len(children["Pod.v1"])
    }

    # Generate the desired child object(s).
    who = parent.get("spec", {}).get("who", "World")
    desired_pods = [
      {
        "apiVersion": "v1",
        "kind": "Pod",
        "metadata": {
          "name": parent["metadata"]["name"]
        },
        "spec": {
          "restartPolicy": "OnFailure",
          "containers": [
            {
              "name": "hello",
              "image": "busybox",
              "command": ["echo", "Hello, %s!" % who]
            }
          ]
        }
      }
    ]

    return {"status": desired_status, "children": desired_pods}

  def do_POST(self):
    # Serve the sync() function as a JSON webhook.
    observed = json.loads(self.rfile.read(int(self.headers.get("content-length"))))
    desired = self.sync(observed["parent"], observed["children"])

    self.send_response(200)
    self.send_header("Content-type", "application/json")
    self.end_headers()
    self.wfile.write(json.dumps(desired).encode())

HTTPServer(("", 80), Controller).serve_forever()
```

Deploy the webhook applying it:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-controller
spec:
  replicas: 1
  selector:
    matchLabels:
      app: hello-controller
  template:
    metadata:
      labels:
        app: hello-controller
    spec:
      containers:
      - name: controller
        image: python:3
        command: ["python3", "/hooks/sync.py"]
        volumeMounts:
        - name: hooks
          mountPath: /hooks
      volumes:
      - name: hooks
        configMap:
          name: hello-controller
---
apiVersion: v1
kind: Service
metadata:
  name: hello-controller
spec:
  selector:
    app: hello-controller
  ports:
  - port: 80
```

Now create and update the custom resource to test the controller:
```yaml
apiVersion: example.com/v1
kind: HelloWorld
metadata:
  name: your-name
spec:
  who: Your Name
```
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


