FROM golang:1.16 as builder
ARG version
ARG commit
ARG date
WORKDIR /src/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o build/ \
    -ldflags="-X sample.version=$version -X sample.commit=$commit -X sample.date=$date" \
    ./...


FROM alpine:3.14
COPY --from=builder /src/build /
ENTRYPOINT ["/sample"]