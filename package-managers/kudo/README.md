# Run the example

Guide [https://kudo.dev/docs/#notes-on-cert-manager](https://kudo.dev/docs/#notes-on-cert-manager)

### 1- Install kudo kubectl plugin and controller 

```bash
kubectl krew install kudo
kubectl kudo init --unsafe-self-signed-webhook-ca
```
Create empty operator scheleton:
```bash
kubectl kudo package new first-operator
```

### 2- Basic test

Check operator inside ./operator folder

Install my operator


## Pros

Interesting declarative approach, no coding required, supports heavy configuration and testing

## Cons

Could not terminate controller upgrade, no easy way of accessing logs, works through kubernetes plugin, not quite usable for production yet

I see it as an extension of helm charts, can get interesting in the future

