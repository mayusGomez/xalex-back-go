GOOS=linux go build getServices.go
zip getServices.zip getServices
aws lambda update-function-code --function-name getServices --zip-file fileb://getServices.zip
rm getServices.zip getServices
