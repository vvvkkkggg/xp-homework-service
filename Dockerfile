# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /src
COPY . .
RUN go build -o /bin/main ./main.go

FROM scratch
COPY --from=0 /bin/main /bin/main
CMD ["/bin/main"]

