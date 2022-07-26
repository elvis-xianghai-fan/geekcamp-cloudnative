FROM alpine
ENV SERVICE_PORT=8080
ADD bin/amd64/httpserver /httpserver
EXPOSE ${SERVICE_PORT}
ENTRYPOINT /httpserver
