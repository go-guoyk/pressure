FROM golang:1.12 AS build
ENV PROJECT pressure
ENV GOPROXY https://goproxy.io
WORKDIR /src/$PROJECT
COPY . .
RUN CGO_ENABLED=0 go install -a -tags netgo -ldflags=-w

FROM scratch
COPY --from=build /go/bin/pressure /pressure
CMD [ "/pressure" ]
