FROM golang:latest

WORKDIR /cosmonova


RUN curl https://get.ignite.com/cli! | bash

COPY . /cosmonova/

RUN ignite chain init && echo done

