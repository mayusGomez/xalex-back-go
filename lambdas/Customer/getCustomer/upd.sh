GOOS=linux go build getCustomer.go
zip getCustomer.zip getCustomer
aws lambda update-function-code --function-name getCustomer --zip-file fileb://getCustomer.zip
rm getCustomer.zip getCustomer