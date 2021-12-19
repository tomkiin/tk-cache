# 获取脚本当前目录
WORK_PATH=$(dirname $0)

# 跳转到脚本目录
cd ./${WORK_PATH}

# 获取当前绝对路径
WORK_PATH=$(pwd)

# 清除正在运行的旧容器
docker rm -f $(docker ps -a | grep "tk-cache" | awk '{print $1}')
echo "\033[32m[1/6] 清理旧容器完成\033[0m"

# 制作 group 镜像
sh ./group/build_image.sh
echo "\033[32m[2/6] 制作 group 镜像完成:\033[0m tk-cache-group"

# 制作 node 镜像
sh ./node/build_image.sh
echo "\033[32m[3/6] 制作 node 镜像完成:\033[0m tk-cache-node"

# 创建互联网络
docker network create -d bridge tk-cache-network
echo "\033[32m[4/6] 创建容器网络完成:\033[0m tk-cache-network"

# 启动 group 容器
docker run -itd --name tk-cache-group --network tk-cache-network -e PARAMS="-replicas=3" -p 8080:8080 tk-cache-group:v1.0
echo "\033[32m[5/6] 创建 group 容器完成\033[0m"

# 等待 group 启动完毕
sleep 3s

# 启动 3 个 node 容器
for i in {1..3}
do
    docker run -itd --name tk-cache-node${i} --network tk-cache-network -e PARAMS="-max_size=10240 -register_ip=tk-cache-group" tk-cache-node:v1.0
done
echo "\033[32m[6/6] 创建 node 容器完成\033[0m"