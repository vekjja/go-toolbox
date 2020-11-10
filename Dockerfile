FROM --platform=${BUILDPLATFORM} golang:1.15-alpine as build

ARG APP
ARG TARGETOS
ARG TARGETARCH

ENV CGO_ENABLED=0
WORKDIR /src

copy common ./common
COPY ${APP} ./${APP}

RUN cd ${APP} &&\
    go mod download &&\
    GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /${APP} .

FROM scratch AS bin
ARG APP
COPY --from=build /${APP} /