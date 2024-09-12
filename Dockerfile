# ---- 
# build executable binary
FROM golang:latest as builder

WORKDIR $GOPATH/src/github.com/osscameroon/
# We only copy our app
COPY . .


WORKDIR $GOPATH/src/github.com/osscameroon/
# We fetch dependencies
RUN go get -d -v

# We compile
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/sammy

# ---
# Let's build our small image
FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/sammy  /go/bin/sammy

# Run the hello binary.
ENTRYPOINT ["/go/bin/sammy"]

