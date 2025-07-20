#!/bin/bash

# 生成 swagger 文件
goctl api swagger --api ./admin/admin.api --dir .

# JSON 文件路径
JSON_FILE="./admin/admin.json"

# 检查文件是否存在
if [ ! -f "$JSON_FILE" ]; then
    echo "错误：JSON 文件 '$JSON_FILE' 不存在。"
    exit 1
fi

# 使用 cat 读取文件内容到字符串变量
# `$(...)` 是命令替换，会将命令的输出赋值给变量
JSON_STRING=$(cat "$JSON_FILE")
echo "读取到的 JSON 字符串："
echo "$JSON_STRING"

# 1. 使用 jq 读取并处理 JSON 文件内容
#    这里不再直接将文件内容读入一个变量。
#    jq . "$JSON_FILE" 会读取文件内容，并以压缩（无额外空格和换行）格式输出原始 JSON
#    这将作为 "input" 字段的值。
#    构建一个 JSON 对象，其中 "input" 的值是来自 admin.json 的内容
#    并保留 "options" 部分。
#    `.` 表示读取整个文件内容
#    @json 表示将一个字符串字面量解析为JSON值
#    注意，这里`--argfile` 和 `@json` 是jq的高级用法，用于将文件内容作为JSON字符串嵌入
REQUEST_BODY=$(jq --null-input \
  --argfile json_input "$JSON_FILE" \
  '{
    "input": ($json_input | tojson),  # tojson将json_input内容序列化为JSON字符串
    "options": {
        "targetEndpointFolderId": 0,
        "targetSchemaFolderId": 0,
        "endpointOverwriteBehavior": "OVERWRITE_EXISTING",
        "schemaOverwriteBehavior": "OVERWRITE_EXISTING",
        "updateFolderOfChangedEndpoint": false,
        "prependBasePath": false
    }
  }'
)


# 同步 API 文档
curl --location -g --request POST 'https://api.apifox.com/v1/projects/4104614/import-openapi?locale=zh-CN' \
--header 'X-Apifox-Api-Version: 2024-03-28' \
--header 'Authorization: Bearer APS-50KsqbUAOaZd86rjFhOLJtzODXykjkk4' \
--header 'Content-Type: application/json' \
--data "$REQUEST_BODY"