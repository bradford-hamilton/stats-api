FROM golang:1.20

# Copy src code into image, fetch dependencies
WORKDIR $GOPATH/src/bradford-hamilton/stats-api
COPY . .
RUN go mod download

# Build for linux 64 bit
ENV STATS_API_ENVIRONMENT="production"
WORKDIR $GOPATH/src/bradford-hamilton/stats-api/cmd/server
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/server .
EXPOSE 4000
ENTRYPOINT ["/go/bin/server"]
