FROM golang:1.22rc1-bullseye

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
ENV PATH="$PATH:$GOROOT/bin:$GOPATH/bin"
ENV PATH="$PATH:$(go env GOPATH)/bin"

WORKDIR /app/
COPY . .

RUN apt-get update \
  && apt-get install -y

# /appで実行しないとエラー出る
# RUN go mod tidy

RUN go install github.com/cosmtrek/air@latest \
  && go install github.com/go-delve/delve/cmd/dlv@latest

# RUN go mod init main \
#   && go mod tidy \
#   && go build

EXPOSE 8080

# CMD ["go", "run", "main.go"]
# CMD ["air", "-c", ".air.toml"]
# ENTRYPOINT cd app && go run main.go
