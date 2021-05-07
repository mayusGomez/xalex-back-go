GOOS=linux go build createQuote.go
zip createQuote.zip createQuote
aws lambda create-function --function-name createQuote --zip-file fileb://createQuote.zip --handler createQuote --runtime go1.x --role arn:aws:iam::$ACCOUNT_ID:role/lambda-basic-rol --environment "Variables={STR_MONGO_CONN=$STR_MONGO_CONN, AUTH0_AUDIENCE=$AUTH0_AUDIENCE, AUTH0_DOMAIN=$AUTH0_DOMAIN}"
rm createQuote.zip createQuote
