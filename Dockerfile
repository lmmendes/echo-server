FROM chainguard/wolfi-base:latest

USER nonroot

# Create necessary directories with correct permissions
WORKDIR /app

# Copy binary
COPY --chown=nonroot:nonroot echo-server .

# Expose ports
EXPOSE 8080

# Run the application
ENTRYPOINT ["./echo-server"]
