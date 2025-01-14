


```shell

// 上传文件
curl -X POST http://127.0.0.1:8000/upload/file \
  -F file=@./storage/uploads/${image} \
  -F type=1

// 创建 tag
curl -X POST http://127.0.0.1:8000/api/v1/tags -F 'name=${name}' -F created_by=${created_by}

// 获取 token
curl -v -X POST 'http://127.0.0.1:8000/auth?app_key=${app_key}&app_secret=${app_secret}'

// token 作为请求头
curl -X GET http://localhost:8080/api/v1/tags \
  -H "token: ${token}"

// token 作为查询参数
curl -X GET "http://localhost:8080/api/v1/tags?token=${token}"

// 限流器测试
for i in {1..11}; do
  curl -X POST 'http://127.0.0.1:8000/auth?app_key=${app_key}&app_secret=${app_secret}'
done

```