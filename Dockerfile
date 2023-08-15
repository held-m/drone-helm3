FROM golang:1.20.6-alpine3.18 AS base

RUN apk update && apk upgrade 

ARG UID=1001
ARG GID=1001
ARG WORKDIR=/app
ARG DRONE_TAG

ENV USR=app
ENV GRP=app
ARG HOME=/home/${USR}

RUN addgroup --gid "${GID}" "${GRP}"
RUN adduser \
    -S \
    -u "${UID}" \
    -g "${GID}" \
    -D "${GRP}" \
   -h ${HOME} \
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

# TODO: use non-root user. Need to fix permision with drone-runner first
# ARG UID=1001
# ARG GID=1001
# ENV USR=app
# ENV GRP=app
# ARG HOME=/home/${USR}
ARG HOME=/root

# TODO: use non-root user. Need to fix permision with drone-runner first
# RUN addgroup --gid "${GID}" "${GRP}"
# RUN adduser \
#     -S \
#     -u "${UID}" \
#     -g "${GID}" \
#     -D "${GRP}" \
#    -h ${HOME} \
#     "${USR}" 

COPY --from=builder /app/drone-helm /bin/drone-helm
COPY assets/kubeconfig.tpl ${HOME}/.kube/config.tpl
# RUN chown -R ${USR}:${GRP} ${HOME} && chmod -R 700 ${HOME}/.kube
RUN chmod -R 700 ${HOME}/.kube

# TODO: use non-root user. Need to fix permision with drone-runner first
# USER ${USR}
WORKDIR ${HOME}


LABEL description="Helm 3 plugin for Drone 3"
LABEL base="alpine/helm"

ENTRYPOINT [ "/bin/drone-helm" ]
