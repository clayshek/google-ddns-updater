FROM arm32v7/golang as builder

WORKDIR /go/src/google-ddns-updater

COPY src/google-ddns-updater.go .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o google-ddns-updater .


FROM arm32v7/alpine

LABEL maintainer="clay@clayshekleton.com"

WORKDIR /root

RUN apk update
#RUN apk upgrade
RUN apk add --no-cache ca-certificates
RUN update-ca-certificates 2>/dev/null || true

COPY --from=builder /go/src/google-ddns-updater/google-ddns-updater .

ENTRYPOINT ["./google-ddns-updater"]

