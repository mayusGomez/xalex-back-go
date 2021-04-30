GOOS=linux go build createCustomer.go
zip createCustomer.zip createCustomer
aws lambda update-function-code --function-name createCustomer --zip-file fileb://createCustomer.zip
rm createCustomer.zip createCustomer