FROM golang
WORKDIR /app
RUN go mod init best-practices-go
RUN apt-get update && apt-get install -y ca-certificates
RUN apt-get install -y tzdata
COPY . .
CMD ["go", "run", "."]