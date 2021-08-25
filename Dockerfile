FROM golang:latest
ENV GOPROXY https://mirrors.aliyun.com/goproxy/
ENV GO111MODULE on
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/release
ADD . .
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o xiuscan-linux main.go
RUN GOOS=windows CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o xiuscan-windows.exe main.go
RUN GOOS=darwin CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o xiuscan-darwin main.go
CMD ["bash","-c","tail -f /dev/null"]
