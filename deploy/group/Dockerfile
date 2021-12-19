FROM centos:7

COPY bin/tk-cache-group /

RUN chmod 777 /tk-cache-group

# 通过环境变量传入参数
ENV PARAMS = ""

ENTRYPOINT ["sh", "-c", "/tk-cache-group $PARAMS"]