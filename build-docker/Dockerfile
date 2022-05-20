FROM golang:1.18-bullseye

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-server .

CMD /app/hello-server 