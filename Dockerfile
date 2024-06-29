FROM golang:1.22

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go/go.mod go/go.sum ./
RUN go mod download && go mod verify

COPY go/. .
RUN go build -v -o /usr/local/bin/app

COPY common/prompt_summary.txt .

CMD ["app"]
