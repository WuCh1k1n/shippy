FROM alpine:latest

RUN mkdir -p /app

WORKDIR /app

COPY vessel-service /app/vessel-service

EXPOSE 50051

CMD [ "./vessel-service" ]