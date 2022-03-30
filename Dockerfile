FROM golang
RUN mkdir -p /server
WORKDIR  /server
COPY . .
EXPOSE 8080
RUN go mod download
RUN go build -o server
ENTRYPOINT ["./server"]