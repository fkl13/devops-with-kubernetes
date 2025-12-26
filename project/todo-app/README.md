# Todo App

Build image:
```
docker image build -t todo-app .
```

Run as a container:
```
docker run -p 8080:8080 -e PORT=8080 todo-app:latest
```

Import image:
```
k3d image import todo-app
```

Apply deployment:
```
kubectl create -f deployment.yaml
```

Check logs:
```
kubectl get pods
kubectl logs <pod-name>
```
