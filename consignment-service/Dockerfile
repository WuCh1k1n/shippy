FROM alpine:latest

RUN mkdir -p /app

WORKDIR /app

COPY consignment-service /app/consignment-service

EXPOSE 50051

CMD [ "./consignment-service" ]