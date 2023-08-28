# build stage
FROM golang:1.19.9-alpine3.17 as backend

ADD ./random-scheduler /bin/

ENTRYPOINT ["/bin/random-scheduler"]