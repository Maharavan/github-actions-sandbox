FROM go:1.25.0-alphine


WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /tool

CMD ["/tool"]

