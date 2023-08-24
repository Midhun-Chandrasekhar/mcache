FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /mcache

EXPOSE 4567

CMD [ "/mcache" ]