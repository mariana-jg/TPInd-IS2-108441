# syntax=docker/dockerfile:1
FROM golang:1.24 AS build

WORKDIR /src
COPY . .  

RUN go mod tidy
RUN go build -o main .

FROM gcr.io/distroless/base-debian12

WORKDIR /root/
COPY --from=build /src/main .

CMD [ "./main" ]

