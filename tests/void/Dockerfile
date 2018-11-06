FROM voidlinux/voidlinux
RUN xbps-install -v -Syu && xbps-install -v -Syu && xbps-install -v -Sy go
RUN go get -u -v github.com/xyproto/distrodetector/cmd/d4
CMD d4
