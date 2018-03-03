FROM centos:centos7

LABEL maintainer="TODO"

ENV MANAGEF_API_HOME=/opt/managef \
    PATH=$MANAGEF_API_HOME:$PATH

COPY managef_api $MANAGEF_API_HOME/


ENTRYPOINT ["/opt/managef/managef_api"]