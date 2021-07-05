FROM golang:1.16

ENV SYMBOL=MSFT \
    API_KEY=M4BNP375Y9LGWZQO \
    NDAYS=10 \
    GIN_MODE=release

WORKDIR /opt/stock-checker
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["stock-checker"]