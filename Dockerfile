## Stage 1

FROM golang:1.19 as builder

# Set destination for COPY
RUN mkdir -p /app
WORKDIR /app

### Setting a proxy for downloading modules
#ENV GOPROXY https://proxy.golang.org,direct

### Copy Go application dependency files
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copies everything from your root directory into /app
COPY . . 

### CGO has to be disabled cross platform builds
### Otherwise the application won't be able to start
ENV CGO_ENABLED=0

# Build
RUN GOOS=linux go build -o ./srvctl

## Stage 2

FROM scratch
WORKDIR /app

COPY --from=builder /app/srvctl .
COPY --from=builder /app/html_templates/ /app/html_templates/
COPY --from=builder /app/srvctl.yml .


EXPOSE 8080

CMD [ "/app/srvctl","-c","/app/srvctl.yml" ]
