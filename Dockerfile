FROM golang:1.18-alpine

RUN apk add --no-cache bash

WORKDIR /app

COPY . .

RUN go build -o password_manager . && go install

CMD ["/bin/bash"]