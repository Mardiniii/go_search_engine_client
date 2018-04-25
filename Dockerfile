FROM iron/go:dev
WORKDIR /app

ENV SRC_DIR=/go/src/github.com/Mardiniii/go_search_engine_client/
COPY . $SRC_DIR
COPY views /app/views

RUN cd $SRC_DIR; go build -o go_search_engine_client; cp go_search_engine_client /app/
ENTRYPOINT ["./go_search_engine_client"]
