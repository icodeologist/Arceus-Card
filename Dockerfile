FROM golang:1.24-alpine


WORKDIR /app

COPY go.* ./ 

RUN go mod download

COPY . .

RUN go build -o arceus main.go

EXPOSE 3000

CMD ["./arceus"]
