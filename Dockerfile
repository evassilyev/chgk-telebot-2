FROM alpine:3.14.3
ARG EXPOSE_PORT
EXPOSE $EXPOSE_PORT
WORKDIR /app
COPY dist/app /app/bot
COPY certificates/key.pem /app/key.pem
COPY certificates/cert.pem /app/cert.pem
RUN chmod +x /app/bot
CMD ["./bot"]
