FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Copies everything from your root directory into /app
COPY . . 
RUN go mod download

# Build
RUN GOOS=linux go build -o /docker-srvctl

EXPOSE 8080

CMD [ "/docker-srvctl","-c","srvctl.yml" ]
