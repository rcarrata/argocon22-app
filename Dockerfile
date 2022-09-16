FROM golang:latest as build
WORKDIR /argocon
COPY . /argocon
COPY ui /argocon/ui
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o ./app .

FROM quay.io/rcarrata/malicious-busybox:hacked
COPY --from=0 /argocon/app .
COPY --from=0 /argocon/ui/ ui
RUN echo "Hacking Dockerfile" >> /tmp/hacked
EXPOSE 8080
CMD ["/app"]