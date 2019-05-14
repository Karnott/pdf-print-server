FROM golang:1.12.5-alpine3.9 AS build

WORKDIR $GOPATH/src/github.com/karnott/pdf-print-server/
ADD . $GOPATH/src/github.com/karnott/pdf-print-server/

RUN apk add -U git

ENV GO111MODULE on
RUN go build -o /go/bin/pdf-print-server main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/pdf-print-server main.go


FROM alpine
WORKDIR /app
RUN apk add -U --no-cache xvfb wkhtmltopdf
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/pdf-print-server /app/
EXPOSE 80
CMD ["/app/pdf-print-server"]
