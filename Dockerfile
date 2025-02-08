From golang:1.23.4 as builder
WORKDIR /app
COPY go.mod go.sum main.go ./
RUN go mod tidy
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o simplerouter .
RUN	CGO_ENABLED=0 go build -ldflags '-s -w' -o simplerouter .

FROM scratch

# Copy the compiled binary into the container.  Crucially, it's statically linked.
COPY --from=builder /app/simplerouter /simplerouter

# Expose the port (this is metadata, but good practice).
EXPOSE 8080

# Run the application.
ENTRYPOINT ["/simplerouter"]
