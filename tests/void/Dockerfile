FROM voidlinux/voidlinux
ENV PROJECT distrodetector
ENV PACKAGES gcc git go
RUN xbps-install -Syu && xbps-install -Syu && xbps-install -Sy $PACKAGES
RUN git clone "https://github.com/xyproto/$PROJECT" "/$PROJECT"
WORKDIR "/$PROJECT"
CMD go test
