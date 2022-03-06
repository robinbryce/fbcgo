FROM golang:1.17 as build

WORKDIR /fbcgo

COPY --from=neomantra/flatbuffers /usr/local/include/flatbuffers c/include/flatbuffers
COPY --from=neomantra/flatbuffers /usr/local/lib/libflat* c/lib/
COPY --from=neomantra/flatbuffers /usr/local/lib/cmake/flatbuffers c/lib/cmake/flatbuffers

COPY go.mod go.sum .
RUN go mod tidy


