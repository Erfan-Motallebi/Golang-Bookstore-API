FROM golang:1-alpine3.5 AS build_base

ENV GO111MODULE=on
RUN apk add --no-cache git
WORKDIR /app/bookstore

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# unit tests 
RUN CGO_ENABLED=0 go test -v

#
RUN go build -o ./out/bookstore .

FROM alpine:3.9.6
RUN apk add ca-certificates

COPY --from=build_base /app/bookstore/out/bookstore /app/bookstore
EXPOSE 8050

CMD [ "/app/bookstore" ]