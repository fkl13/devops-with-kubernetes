FROM golang:1.25 AS build

WORKDIR /src/
COPY go.mod ./
# RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/todo-app

FROM scratch
COPY --from=build /bin/todo-app /bin/todo-app
ENTRYPOINT ["/bin/todo-app"]
