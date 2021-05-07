GOOS=linux go build getQuotes.go
zip getQuotes.zip getQuotes
aws lambda update-function-code --function-name getQuotes --zip-file fileb://getQuotes.zip
rm getQuotes.zip getQuotes