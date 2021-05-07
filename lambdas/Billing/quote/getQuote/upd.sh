GOOS=linux go build quote.go
zip quote.zip quote
aws lambda update-function-code --function-name quote --zip-file fileb://quote.zip
rm quote.zip quote