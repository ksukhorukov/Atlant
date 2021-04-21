FROM golang:latest AS build
RUN mkdir -p /usr/src/atlant
WORKDIR /usr/src/atlant
COPY ./server/server ./server

FROM scratch
COPY --from=build /usr/src/atlant/server /bin/server
ENTRYPOINT /usr/bin/echo
