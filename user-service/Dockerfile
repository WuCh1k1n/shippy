FROM alpine:latest

RUN mkdir -p /app

WORKDIR /app

COPY user-service /app/user-service

EXPOSE 50051

CMD [ "./user-service" ]