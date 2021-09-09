# --build
FROM golang:1.16-alpine AS build

WORKDIR /app 

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /new-bank-api

# --deploy
FROM gcr.io/distroless/base-debian10

WORKDIR / 

COPY --from=build /new-bank-api /new-bank-api

EXPOSE $PORT

USER nonroot:nonroot

ENTRYPOINT ["/new-bank-api"]