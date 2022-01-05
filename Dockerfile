FROM golang:1.17 as builder
RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get -y install upx
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd cmd && go build -ldflags="-w -s" -o ../myapp
RUN upx --best --lzma myapp

FROM scratch
WORKDIR /app
COPY --from=builder /app/myapp /app
ENTRYPOINT ["/app/myapp"]