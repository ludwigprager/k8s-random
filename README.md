# k8s Toolbox
This is a toolbox for various testing and troubleshooting on Kubernetes.   
I use it on kubernetes to verify networking et.al.

# Run In Kubernetes
kubectl run -ti --rm toolbox-$RANDOM --image=ludwigprager/k8s-random:latest

# Run In Docker
```bash
docker run -ti \
  --name toolbox \
  ludwigprager/k8s-random:latest \
  /bin/bash
```

or

docker exec -ti toolbox /bin/sh

# Build The Image Yourself
```bash
source set-env.sh
docker build -t $K8S_RANDOM_IMAGE .
```
