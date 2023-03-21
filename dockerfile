FROM golang:latest

ARG EXPOSE_PORT

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o chatericu

EXPOSE ${EXPOSE_PORT}

CMD ["./chatericu"]