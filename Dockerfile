FROM golang:1.17

RUN apt-get update && apt-get install build-essential -y

# Copy code
RUN mkdir -p /code
COPY . /code/
WORKDIR /code

RUN make install
RUN make build

# RUN go mod tidy && CGO_ENABLED=0 go build && ./go-blockchain