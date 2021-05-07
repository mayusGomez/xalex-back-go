GOOS=linux go build createQuote.go
zip createQuote.zip createQuote
aws lambda update-function-code --function-name createQuote --zip-file fileb://createQuote.zip
rm createQuote.zip createQuote