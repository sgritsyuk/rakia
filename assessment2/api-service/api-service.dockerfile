FROM alpine:latest

RUN mkdir /app

COPY api_bin /app

CMD [ "/app/api_bin" ]