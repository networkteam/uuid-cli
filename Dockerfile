FROM golang:1.13 AS build

WORKDIR /app

# Optimization to cache dependencies
ADD go.* ./
RUN go mod download

ADD . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o uuid-cli ./cmd

FROM scratch
COPY --from=build /app/uuid-cli /
ENTRYPOINT ["/uuid-cli"]
