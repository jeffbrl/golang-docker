FROM golang:alpine3.13

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o server . 

# Run the server executable
CMD [ "/app/server" ]