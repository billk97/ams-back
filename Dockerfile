FROM golang:1.18-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN ls -la
RUN go build -o ams-back
CMD [ "./ams-back" ]
