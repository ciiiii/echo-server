FROM golang:1.13-alpine
WORKDIR /work
ADD . .
ENV GOPROXY https://goproxy.io
RUN go build -o /bin/grpcServer .
WORKDIR /
RUN rm -rf /work
EXPOSE 8000
CMD ["/bin/grpcServer"]
