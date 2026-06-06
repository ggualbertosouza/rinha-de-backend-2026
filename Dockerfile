FROM golang:1.25.5 AS build

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOAMD64=v3

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath -ldflags="-s -w" -o /out/server ./cmd/server
RUN go build -trimpath -ldflags="-s -w" -o /out/lb ./cmd/lb

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /out/server /server
COPY --from=build /out/lb /lb

COPY index/ /index/