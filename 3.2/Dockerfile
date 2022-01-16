FROM ubuntu

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y curl

ADD /demo/ /demo

EXPOSE 3000

ENTRYPOINT /demo