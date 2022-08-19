FROM golang:1.16-alpine as build-stage


RUN apk add --no-cache

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN GOOS=linux go build -o s3gallery /app/cmd


FROM alpine
COPY --from=build-stage /app/s3gallery /
EXPOSE 8000
ENTRYPOINT ["/s3gallery"]