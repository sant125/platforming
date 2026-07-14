from golang:1.22 AS builder
workdir /app
copy src/go.mod src/go.sum ./
run go mod download
copy src/ .
run CGO_ENABLED=0 GOOS=linux go build -o api .

from scratch 
copy --from=builder /app/api .
USER 65532:65532
expose 8080
ENTRYPOINT ["./api"]
