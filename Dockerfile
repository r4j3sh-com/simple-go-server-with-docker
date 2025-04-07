# Simple-Go-Server-with-docker

# For environment using lightweight Debian Linus OS
FROM debian:stable-slim

# Copy source to destination
COPY simple-go-server-with-docker /bin/simple-go-server-with-docker

# Automatically start the server
CMD ["/bin/simple-go-server-with-docker"]


