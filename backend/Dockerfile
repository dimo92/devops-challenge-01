# build stage
FROM golang:1.16-alpine AS builder
COPY . /app
WORKDIR /app
RUN ls -la
RUN go mod download 

RUN go get -d -v 
RUN go build -o bin/api 

RUN ls -la bin/

FROM alpine
WORKDIR /app
COPY --from=builder /app/bin/api /bin/
ENTRYPOINT ["/bin/api"]