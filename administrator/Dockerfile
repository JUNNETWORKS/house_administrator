FROM golang:1.14-buster AS build

WORKDIR /
COPY . /src
RUN go get github.com/gin-gonic/gin
RUN go get google.golang.org/grpc
RUN go get github.com/golang/protobuf/ptypes/empty
RUN cd src/ && go build -o administrator main.go


FROM alpine:3.7

COPY --from=build /src/administrator /usr/local/bin/
CMD [ "administrator" ]