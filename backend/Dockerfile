FROM golang:1.19.3-alpine3.17

# Define current working directory
WORKDIR /bucks

# Download modules to local cache so we can skip re-
# downloading on consecutive docker build commands
COPY go.mod .
COPY go.sum .
RUN go mod download

# Add sources
COPY . .

RUN go build -o out/auth0-go-web-app .

# Expose port 3000 for our web app binary
EXPOSE 3000

CMD ["/bucks/out/auth0-go-web-app"]
