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
