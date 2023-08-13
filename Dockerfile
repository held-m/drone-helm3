FROM golang:1.20.6-alpine3.18 AS base

RUN apk update && apk upgrade 

ARG UID=1001
ARG GID=1001
ARG HOME=/home/app
ARG WORKDIR=/app
ARG DRONE_TAG

ENV USR=app
ENV GRP=app

RUN addgroup --gid "${GID}" "${GRP}"
RUN adduser \
    -S \
    -u "${UID}" \
    -g "${GID}" \
    -D "${GRP}" \
   -h /home/app \
    "${USR}" 

RUN mkdir ${WORKDIR}
RUN chown -R ${USR}:${GRP} ${WORKDIR}

USER ${USR}

WORKDIR ${WORKDIR}

FROM base AS builder

ARG DRONE_TAG

COPY . /app/
RUN go clean -modcache \
    && go build -o drone-helm \
    ./cmd/drone-helm/main.go

FROM alpine/helm:3.12.3
MAINTAINER Joachim Hill-Grannec <joachim@pelo.tech>

# COPY build/drone-helm /bin/drone-helm
COPY --from=builder /app/drone-helm /bin/drone-helm
COPY assets/kubeconfig.tpl /root/.kube/config.tpl

LABEL description="Helm 3 plugin for Drone 3"
LABEL base="alpine/helm"

ENTRYPOINT [ "/bin/drone-helm" ]
