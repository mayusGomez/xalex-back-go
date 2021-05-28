GOOS=linux go build getUser.go
zip getUser.zip getUser
aws lambda update-function-code --function-name getUser --zip-file fileb://getUser.zip
rm getUser.zip getUser