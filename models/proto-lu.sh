#!/bin/sh

# 获取当前工作目录
CURRENT_DIR=$(pwd)

# 定义 proto 和 pb 目录的相对路径
PROTO_DIR="$CURRENT_DIR/proto"
PB_DIR="$CURRENT_DIR/pb"

# 创建 pb 目录（如果不存在）
mkdir -p "$PB_DIR"

# 检查 proto 目录是否存在
if [ ! -d "$PROTO_DIR" ]; then
    echo "Error: Directory $PROTO_DIR does not exist."
    exit 1
fi

# 遍历 proto 目录下的所有 .proto 文件
for proto_file in "$PROTO_DIR"/*.proto ; do
    echo "$proto_file"
    # 生成对应的 pb 文件
    protoc --proto_path="$PROTO_DIR" \
           --go_out="$PB_DIR" \
           --go-grpc_out="$PB_DIR" \
           "$proto_file"
done