# build stage
FROM golang as builder

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"

WORKDIR /app
RUN mkdir /app/apks

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o apkbackend .

# final stage
FROM scratch
COPY --from=builder /app/apkbackend .
COPY --from=builder /app/.env .
EXPOSE 8686
ENTRYPOINT ["/apkbackend"]
