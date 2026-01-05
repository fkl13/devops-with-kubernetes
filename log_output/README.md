# Log output app

Build image:
```
docker image build -t "logoutput:1.0" .
```

Run as a container:
```
docker run logoutput:1.0
```

Import image:
```
k3d image import logoutput:1.0
```

Apply deployment:
```
kubectl apply -f manifests/deployment.yaml
```

Check logs:
```
kubectl get pods
kubectl logs <pod-name>
