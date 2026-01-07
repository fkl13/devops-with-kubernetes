# Todo App

Build image:
```
docker image build -t fkl13/todo-app:1.0 .
```

Run as a container:
```
docker run -p 8080:8080 -e PORT=8080 fkl13/todo-app:1.0
```

Apply deployment:
```
kubectl apply -f deployment.yaml
```

Check logs:
```
kubectl get pods
kubectl logs <pod-name>
```

## 1.5

Build image:
```
docker image build -t fkl13/todo-app:1.1 .
```

Run as a container:
```
docker run -e PORT=3000 -p 3000:3000 fkl13/todo-app:1.1
```

Apply deployment:
```
kubectl apply -f deployment.yaml
```

Send request:
```
curl localhost:3000
```
