FROM alpine
ADD LPS-srv /LPS-srv
ENTRYPOINT [ "/LPS-srv" ]
