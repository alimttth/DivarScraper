FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    wget \
    ca-certificates

RUN wget https://golang.org/dl/go1.17.7.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.17.7.linux-amd64.tar.gz && \
    rm go1.17.7.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

ENV GOPATH=/go

WORKDIR /go/src/app
COPY . .

RUN go build -o main .

CMD ["./main"]
