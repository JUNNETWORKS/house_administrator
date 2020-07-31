FROM golang:1.14-alpine AS build

COPY . /go/src/github.com/JUNNETWORKS/house_administrator
WORKDIR /go/src/github.com/JUNNETWORKS/house_administrator
RUN go build -o administrator main.go

FROM alpine:3.12

COPY --from=build /go/src/github.com/JUNNETWORKS/house_administrator /usr/local/bin/
CMD [ "administrator" ]
