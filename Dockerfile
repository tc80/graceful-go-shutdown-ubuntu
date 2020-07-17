FROM debian:stretch

# avoid selecting location for tzdata
ARG DEBIAN_FRONTEND=noninteractive

# maybe map file system -> docker system?
# run with no update or fake it ?

RUN apt update
RUN apt-get install -y golang-go=1.13
RUN apt install -y git
RUN apt-get install -y vim
RUN apt-get install -y systemd
RUN apt-get install make

# RUN tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
# RUN export PATH=$PATH:/usr/local/go/bin

ENV GOPATH=/usr/go
WORKDIR /usr/go/src/github.com/cdnjs
RUN git clone https://github.com/cdnjs/tools

WORKDIR /usr/go/src/github.com/cdnjs/tools
# RUN make
RUN echo 'export tools="/usr/go/src/github.com/cdnjs/tools"' >> ~/.bashrc

WORKDIR /usr/go/src
COPY . .
