GOOS=linux go build createService.go
zip createService.zip createService
aws lambda update-function-code --function-name createService --zip-file fileb://createService.zip
rm createService.zip createService