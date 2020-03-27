FROM golang:latest as builder

RUN mkdir /ProjectHandler
WORKDIR /ProjectHandler
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .

FROM scratch
COPY --from=builder /ProjectHandler .

COPY www /static

ENTRYPOINT [ "./ProjectHandler" ]