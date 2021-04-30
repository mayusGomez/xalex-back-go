GOOS=linux go build getCustomers.go
zip getCustomers.zip getCustomers
aws lambda update-function-code --function-name getCustomers --zip-file fileb://getCustomers.zip
rm getCustomers.zip getCustomers