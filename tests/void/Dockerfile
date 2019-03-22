FROM voidlinux/voidlinux
ENV PROJECT distrodetector
ENV PACKAGES gcc git go vim nano bash st-terminfo unibilium ncurses-base
RUN xbps-install -Syu && xbps-install -Syu && xbps-install -Sy $PACKAGES
RUN git clone "https://github.com/xyproto/$PROJECT" "/$PROJECT"
WORKDIR "/$PROJECT"
ENV TERM xterm
CMD go test
