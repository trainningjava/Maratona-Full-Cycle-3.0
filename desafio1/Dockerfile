FROM golang:1.14.4-alpine

ENV PORT 3000
WORKDIR /usr/app

RUN echo "C O P Y N G . . . . "
COPY . .

RUN echo "=================="
RUN echo "B U I L D . . . . "
RUN echo "=================="

RUN go build main.go
RUN rm main.go

ENV ADDR=0.0.0.0
EXPOSE 3000

CMD  ["./main"]