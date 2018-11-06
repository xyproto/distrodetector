FROM voidlinux/voidlinux
ENV PACKAGES git go
RUN xbps-install -Syu && xbps-install -Syu && xbps-install -Sy "$PACKAGES"
RUN go get -u -v github.com/xyproto/distrodetector/cmd/d4
CMD d4
