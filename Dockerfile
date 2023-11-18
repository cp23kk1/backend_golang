FROM golang:1.18

WORKDIR /go/src/app

# ARG DB_USER
# ARG DB_PASSWORD
ARG DB_HOST
# ARG DB_PORT
# ARG DB_NAME

# ENV DB_USER=${DB_USER}
# ENV DB_PASSWORD=${DB_PASSWORD}
# ENV DB_NAME=${DB_NAME}
ENV DB_HOST=${DB_HOST}
# ENV DB_PORT=${DB_PORT}

COPY . .

RUN go mod tidy
RUN go mod download
RUN go build -o main main.go

CMD ["./main"]

EXPOSE 8080