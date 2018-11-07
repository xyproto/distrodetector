FROM debian
ENV PROJECT distrodetector
ENV GOVER 1.11.2
ENV PACKAGES curl git
ENV PATH /go/bin:$PATH
RUN apt-get update && apt-get -y upgrade && apt-get -y install $PACKAGES
RUN curl -sOL "https://dl.google.com/go/go$GOVER.linux-amd64.tar.gz"
RUN tar -C / xvf "go$GOVER.linux-amd64.tar.gz"
RUN git clone "https://github.com/xyproto/$PROJECT" "/$PROJECT"
WORKDIR "/$PROJECT"
CMD go test
