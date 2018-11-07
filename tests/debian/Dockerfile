FROM debian
ENV PROJECT distrodetector
ENV PACKAGES git golang
RUN apt-get -y update && apt-get -y upgrade && apt-get -y install $PACKAGES
RUN git clone "https://github.com/xyproto/$PROJECT" "/$PROJECT"
WORKDIR "/$PROJECT"
CMD go test
