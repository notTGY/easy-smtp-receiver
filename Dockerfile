FROM ghcr.io/nottgy/easy-smtp-receiver
ENTRYPOINT ["/bin/sh", "-c", "cd /app && ./server"]
