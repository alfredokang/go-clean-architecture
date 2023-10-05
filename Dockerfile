FROM golang:1.20.4
LABEL authors="alfredokang"
WORKDIR /usr/app
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0

CMD ["tail", "-f", "/dev/null"]

ENTRYPOINT ["top", "-b"]