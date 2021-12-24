# 测试在高并发环境下，singleflight 机制是否生效
CONCURRENT=$1

curl -XPOST -d '{"key": "testKey", "value": "testValue"}' http://127.0.0.1:8080/cache

for i in $(seq 1 ${CONCURRENT})
do
    curl http://127.0.0.1:8080/cache\?key=testKey &
done

wait