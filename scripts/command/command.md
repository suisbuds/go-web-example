


```shell
curl -X POST http://127.0.0.1:8000/upload/file \
  -F file=@./storage/uploads/yorushika.png \
  -F type=1


curl -X POST http://127.0.0.1:8000/api/v1/tags -F 'name=doppler' -F created_by=suisbuds
```