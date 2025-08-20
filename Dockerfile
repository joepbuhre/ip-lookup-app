FROM node:20-bullseye-slim AS builder

WORKDIR /app

COPY ./frontend .

RUN npm ci
RUN npm run build

FROM golang:1.23-alpine AS gobuilder

WORKDIR /app

COPY --from=builder /app/dist ./frontend/dist

COPY ./*.go .
COPY ./go.mod go.mod

RUN CGO_ENABLED=0 go build -tags embed -o ./ip-lookup-app

FROM scratch

WORKDIR /app

COPY --from=gobuilder /app/ip-lookup-app .

ENTRYPOINT [ "./ip-lookup-app" ]