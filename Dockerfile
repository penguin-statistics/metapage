FROM node:16 AS webbuilder

WORKDIR /build

COPY . .

# RUN apt -y update && apt -y install git make gcc g++ yarn

RUN make buildweb

FROM golang:1.18-alpine AS base
WORKDIR /app

# builder
FROM base AS gobuilder
ENV GOOS linux
ENV GOARCH amd64

# modules: utilize build cache
COPY go.mod ./
COPY go.sum ./

# RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .

RUN apk add git

# build the binary
RUN go build -o metapagebackend .

# runner
FROM base AS runner
RUN apk add --no-cache libc6-compat

RUN apk add --no-cache tini
# Tini is now available at /sbin/tini

COPY --from=gobuilder /app/metapagebackend /app/metapagebackend

COPY --from=webbuilder /build/web/dist /app/web/dist

ENTRYPOINT ["/sbin/tini", "--"]
CMD [ "/app/metapagebackend" ]
