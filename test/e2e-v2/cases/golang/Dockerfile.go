# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.18 AS builder

ARG SW_AGENT_GOLANG_COMMIT
ARG GOLANG_CODE=${SW_AGENT_GOLANG_COMMIT}.tar.gz
ARG GOLANG_CODE_URL=https://github.com/apache/skywalking-go/archive/${GOLANG_CODE}

ENV CGO_ENABLED=0
ENV GO111MODULE=on

WORKDIR /

RUN ls -a

WORKDIR /skywalking-go

ADD ${GOLANG_CODE_URL} .
RUN tar -xf ${GOLANG_CODE} --strip 1
RUN rm ${GOLANG_CODE}

RUN ls -a

COPY ./main.go /e2e/main.go
COPY ./go.mod /e2e/go.mod
COPY ./go.sum /e2e/go.sum

WORKDIR /e2e

RUN ls -a

RUN go build -o main

ENTRYPOINT ["./main"]
