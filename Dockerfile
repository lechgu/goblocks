FROM golang:alpine as build

WORKDIR /app

COPY . .
RUN rm -rf cmd/server/html \
    && mkdir cmd/server/html \
    && cp assets/index.html cmd/server/html/ \
    && cp /usr/local/go/misc/wasm/wasm_exec.js cmd/server/html \
    && GOOS=js GOARCH=wasm go build -o cmd/server/html/main.wasm cmd/desktop/main.go \
    && CGO_ENABLED=0 GOOS=linux go build -o goblocks cmd/server/main.go

FROM scratch  
COPY --from=build /app/goblocks /bin/
ENTRYPOINT [ "/bin/goblocks" ]