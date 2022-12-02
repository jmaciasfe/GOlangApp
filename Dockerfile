FROM golang:1.19-alpine

WORKDIR /GOlangApp

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /appgolang

EXPOSE 8081

CMD [ "/appgolang" ]
