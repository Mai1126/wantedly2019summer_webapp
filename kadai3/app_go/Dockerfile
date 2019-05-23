FROM golang:1.10.2

WORKDIR /go/src/app

RUN go get "github.com/lib/pq"
RUN go get "github.com/gorilla/mux"

CMD ["go", "run", "hello.go"]