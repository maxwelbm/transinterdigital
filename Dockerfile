FROM golang:1.19

# set working directory
WORKDIR /app

ENV PATH="/go/bin:${PATH}"

COPY . .
RUN go get -d -v ./...

# Build the Go app
RUN go build -o api .

EXPOSE 3000

# Run the executable
CMD ["./api"]