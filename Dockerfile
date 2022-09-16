FROM golang:1.18 as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

# Copy the go source
COPY main.go main.go
COPY database/ database/
COPY models/ models/
COPY routers/ routers/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go


FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532
EXPOSE 3000

ENTRYPOINT ["/manager"]