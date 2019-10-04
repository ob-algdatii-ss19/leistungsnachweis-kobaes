FROM golang:latest as build
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app

RUN go get github.com/ogier/pflag
RUN go build -o main .

FROM iron/go
RUN mkdir /app
COPY --from=build /app/main /app/
CMD ["/app/main"]
