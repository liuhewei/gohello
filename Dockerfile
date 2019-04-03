FROM busybox
COPY ./gohello /gohello
CMD ["/gohello"]
