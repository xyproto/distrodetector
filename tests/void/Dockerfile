FROM voidlinux/voidlinux
ENV DISTRODETECT "Linux (void)"
ENV PACKAGES git go
RUN xbps-install -Syu && xbps-install -Syu && xbps-install -Sy $PACKAGES
RUN git clone 'https://github.com/xyproto/distrodetector'
CMD cd distrodetector && go test
