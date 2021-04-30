GOOS=linux go build updateCustomer.go
zip updateCustomer.zip updateCustomer
aws lambda update-function-code --function-name updateCustomer --zip-file fileb://updateCustomer.zip
rm updateCustomer.zip updateCustomer