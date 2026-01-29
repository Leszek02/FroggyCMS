FROM golang:1.25 AS build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o prestaprest -v ./cmd


FROM alpine:3.14 AS main-cms-image

WORKDIR /app
RUN mkdir binary

COPY --from=build /build/prestaprest ./binary

COPY ./public ./public
COPY ./template ./template
COPY ./media ./media

EXPOSE 8080
WORKDIR /app/binary

CMD ["./prestaprest"]