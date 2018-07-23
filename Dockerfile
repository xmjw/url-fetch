FROM iron/go:dev
WORKDIR /app
ENV SRC_DIR=/go/src/github.com/xmjw/url-fetch/

# Add the source code:
ADD . $SRC_DIR

# Build it:
RUN cd $SRC_DIR; go build -o app; cp app /app/

EXPOSE 8080

ENTRYPOINT ["./app"]
