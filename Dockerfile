FROM golang:1.20.5 AS build
WORKDIR /go/src/whisperweb-backend
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o whisperweb-backend .

FROM golang:1.20.5 as release
WORKDIR /app
COPY --from=build /go/src/whisperweb-backend/whisperweb-backend .
EXPOSE 5000
CMD [ "./whisperweb-backend" ]