FROM voidlinux/voidlinux
ENV PACKAGES git go
RUN xbps-install -Syu && xbps-install -Syu && xbps-install -Sy $PACKAGES
RUN git clone github.com/xyproto/distrodetector
CMD DISTRODETECT="Linux (void)" go test
