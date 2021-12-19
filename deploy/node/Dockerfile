FROM centos:7

COPY bin/tk-cache-node /

RUN chmod 777 /tk-cache-node

# 通过环境变量传入参数
ENV PARAMS = ""

ENTRYPOINT ["sh", "-c", "/tk-cache-node $PARAMS"]