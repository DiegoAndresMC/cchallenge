FROM golang:1.17 as build

WORKDIR /app

COPY . .

FROM ubuntu:20.04
RUN apt-get update && apt-get install -y build-essential
RUN make build


FROM golang:1.17 as run
WORKDIR /app
COPY --from=build /app/build/bin/searching-products-wallmart-api /app


ENTRYPOINT [ "/app/build/bin/searching-products-wallmart-api" ][]