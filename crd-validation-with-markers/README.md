# crd-validation-with-markers

## Setup

### Scaffold custom CRDs

```bash
kubebuilder init --domain rewanthtammana.com --license none --owner "rewanthtammana" --plugins=go/v4-alpha
kubebuilder create api --version v1 --group validate --kind ImmutableKind
```

### Create custom CRDs

```bash
make manifests
make install
```

### Deploy sample CRD

```bash
kubectl apply -f ./config/samples/validate_v1_immutablekind.yaml
```

Edit `immutablekind-sample`, remove all labels & try to deploy it. For example,

```bash
echo "apiVersion: validate.rewanthtammana.com/v1
kind: ImmutableKind
metadata:
  name: immutablekind-sample" | kubectl apply -f-
```
