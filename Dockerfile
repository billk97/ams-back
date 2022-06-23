FROM golang:1.18-alpine
WORKDIR /go/app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY * ./
RUN CGO_ENABLED=0 go build -o ams-back
CMD [ "/" ]ams-back