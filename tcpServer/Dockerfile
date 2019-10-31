FROM golang:1.13-alpine
WORKDIR /work
ADD . .
ENV GOPROXY https://goproxy.io
RUN go build -o /bin/tcpServer .
WORKDIR /
RUN rm -rf /work
EXPOSE 3333
CMD ["/bin/tcpServer"]
