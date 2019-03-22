FROM archlinux/base
ENV PROJECT distrodetector
ENV PACKAGES git go
RUN pacman -Syu --noconfirm && pacman -S --noconfirm $PACKAGES
RUN git clone "https://github.com/xyproto/$PROJECT" "/$PROJECT"
WORKDIR "/$PROJECT"
CMD go test
