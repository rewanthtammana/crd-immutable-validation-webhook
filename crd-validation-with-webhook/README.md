# crd-validation-with-webhook

## Setup

### Install cert manager for certificates

```bash
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.9.1/cert-manager.yaml
```

### Scaffold custom CRDs

```bash
kubebuilder init --domain rewanthtammana.com --license none --owner "rewanthtammana" --plugins=go/v4-alpha
kubebuilder create api --version v1 --group validate --kind ImmutableKind
kubebuilder create webhook --group validate --version v1 --kind ImmutableKind --programmatic-validation
```

### Customize the code to enable webhook

* Uncomment `patches/webhook_in_immutablekinds.yaml` and `patches/cainjection_in_immutablekinds.yaml` in `config/crd/kustomization.yaml`
* Uncomment `../certmanager` and `../webhook` directories in `config/default/kustomization.yaml`
* Uncomment `manager_webhook_patch.yaml` in `config/default/kustomization.yaml`
* Uncomment entire `CERTMANAGER` replacements block in `config/default/kustomization.yaml`

### Create custom CRDs & deploy the webhook

```bash
make manifests
make docker-build docker-push IMG=rewanthtammana/immutablekindwebhook:v1
make install deploy IMG=rewanthtammana/immutablekindwebhook:v1
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
