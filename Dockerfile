FROM golang:1.21-alpine AS backend
WORKDIR /app
ENV CGO_ENABLED=0
COPY go.mod go.sum ./
RUN --mount=type=ssh go mod download && go mod verify
COPY . .
RUN go build -o /app/docker-file-management /app/cmd/main/main.go

FROM alpine
COPY --from=backend /app/docker-file-management .
USER 1000
ENTRYPOINT ["/docker-file-management"]
