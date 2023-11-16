# syntax = docker/dockerfile:1.4

ARG TARGET=enduro
ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine AS build-go
WORKDIR /src
ENV CGO_ENABLED=0
COPY --link go.* ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY --link . .

FROM build-go AS build-enduro
ARG VERSION_PATH
ARG VERSION_LONG
ARG VERSION_SHORT
ARG VERSION_GIT_HASH
RUN --mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	go build \
	-trimpath \
	-ldflags="-X '${VERSION_PATH}.Long=${VERSION_LONG}' -X '${VERSION_PATH}.Short=${VERSION_SHORT}' -X '${VERSION_PATH}.GitCommit=${VERSION_GIT_HASH}'" \
	-o /out/enduro .

FROM build-go AS build-enduro-a3m-worker
ARG VERSION_PATH
ARG VERSION_LONG
ARG VERSION_SHORT
ARG VERSION_GIT_HASH
RUN --mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	go build \
	-trimpath \
	-ldflags="-X '${VERSION_PATH}.Long=${VERSION_LONG}' -X '${VERSION_PATH}.Short=${VERSION_SHORT}' -X '${VERSION_PATH}.GitCommit=${VERSION_GIT_HASH}'" \
	-o /out/enduro-a3m-worker \
	./cmd/enduro-a3m-worker

FROM build-go AS build-enduro-am-worker
ARG VERSION_PATH
ARG VERSION_LONG
ARG VERSION_SHORT
ARG VERSION_GIT_HASH
RUN --mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	go build \
	-trimpath \
	-ldflags="-X '${VERSION_PATH}.Long=${VERSION_LONG}' -X '${VERSION_PATH}.Short=${VERSION_SHORT}' -X '${VERSION_PATH}.GitCommit=${VERSION_GIT_HASH}'" \
	-o /out/enduro-am-worker \
	./cmd/enduro-am-worker

FROM alpine:3.18.2 AS base
ARG USER_ID=1000
ARG GROUP_ID=1000
RUN addgroup -g ${GROUP_ID} -S enduro
RUN adduser -u ${USER_ID} -S -D enduro enduro
USER enduro

FROM base AS enduro
COPY --from=build-enduro --link /out/enduro /home/enduro/bin/enduro
COPY --from=build-enduro --link /src/enduro.toml /home/enduro/.config/enduro.toml
CMD ["/home/enduro/bin/enduro", "--config", "/home/enduro/.config/enduro.toml"]

FROM base AS enduro-a3m-worker
COPY --from=build-enduro-a3m-worker --link /out/enduro-a3m-worker /home/enduro/bin/enduro-a3m-worker
COPY --from=build-enduro-a3m-worker --link /src/enduro.toml /home/enduro/.config/enduro.toml
CMD ["/home/enduro/bin/enduro-a3m-worker", "--config", "/home/enduro/.config/enduro.toml"]

FROM base AS enduro-am-worker
# Install python/pip
ENV PYTHONUNBUFFERED=1
USER root
RUN apk add --update --no-cache python3 && \
	ln -sf python3 /usr/bin/python && \
	python3 -m ensurepip
USER enduro
RUN pip3 install --no-cache --upgrade pip lxml bagit==v1.8.1
COPY --from=build-enduro-am-worker --link /out/enduro-am-worker /home/enduro/bin/enduro-am-worker
COPY --from=build-enduro-am-worker --link /src/enduro.toml /home/enduro/.config/enduro.toml
# SFA metadata schema for validation.
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/xsdval.py xsdval.py
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/repackage_sip.py repackage_sip.py
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/ablieferung.xsd ablieferung.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/archivischeNotiz.xsd archivischeNotiz.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/archivischerVorgang.xsd archivischerVorgang.xsd 
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/arelda.xsd arelda.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/base.xsd base.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/datei.xsd datei.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/dokument.xsd dokument.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/dossier.xsd dossier.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/ordner.xsd ordner.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/ordnungssystem.xsd ordnungssystem.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/ordnungssystemposition.xsd ordnungssystemposition.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/paket.xsd paket.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/provenienz.xsd provenienz.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/zusatzDaten.xsd zusatzDaten.xsd
COPY --from=build-enduro-am-worker --link /src/hack/sampledata/xsd/bagit.txt bagit.txt
CMD ["/home/enduro/bin/enduro-am-worker", "--config", "/home/enduro/.config/enduro.toml"]

FROM ${TARGET}
