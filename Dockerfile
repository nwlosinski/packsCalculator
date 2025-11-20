FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

# copy source code
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go


# ---- RUNTIME IMAGE ----
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=build /app/server /app/server
COPY --from=build /app/ui /app/ui
COPY config.json /app/config.json

EXPOSE 80

ENTRYPOINT ["/app/server"]