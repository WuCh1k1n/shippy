FROM alpine:latest

RUN mkdir -p /app

WORKDIR /app

COPY user-cli /app/user-cli

EXPOSE 50051

CMD [ "./user-cli" ]