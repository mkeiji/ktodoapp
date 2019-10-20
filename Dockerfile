FROM golang
WORKDIR /todoApp

COPY go.mod /todoApp/go.mod
COPY go.sum /todoApp/go.sum

RUN go mod download

COPY . /todoApp
ENTRYPOINT ["go", "run", "main.go"]
