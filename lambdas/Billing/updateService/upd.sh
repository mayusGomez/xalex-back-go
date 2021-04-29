GOOS=linux go build updateService.go
zip updateService.zip updateService
aws lambda update-function-code --function-name updateService --zip-file fileb://updateService.zip
rm updateService.zip updateService