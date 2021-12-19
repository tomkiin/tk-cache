# 容器打包信息
IMAGE_NAME="tk-cache-node"
IMAGE_TAG="v1.0"

# 获取脚本当前目录
WORK_PATH=$(dirname $0)

# 跳转到脚本目录
cd ./${WORK_PATH}

# 获取当前绝对路径
WORK_PATH=$(pwd)

# 添加二进制文件临时目录
mkdir bin

# 设置 go 环境变量
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct

# 交叉编译 mac -> linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${WORK_PATH}/bin/ ${WORK_PATH}/../../cmd/tk-cache-node/...
# go build -o ${WORK_PATH}/bin/ ${WORK_PATH}/../

# 制作容器
docker build -t ${IMAGE_NAME}:${IMAGE_TAG} -f ./Dockerfile .

# 清理临时文件夹和 none 镜像
docker image prune -f
rm -rf bin