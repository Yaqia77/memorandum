FROM golang:1.18-alpine

WORKDIR /apps

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o apps

EXPOSE 8080

CMD [ "./apps" ]
