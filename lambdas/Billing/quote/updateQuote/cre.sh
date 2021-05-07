GOOS=linux go build updateQuote.go
zip updateQuote.zip updateQuote
aws lambda create-function --function-name updateQuote --zip-file fileb://updateQuote.zip --handler updateQuote --runtime go1.x --role arn:aws:iam::$ACCOUNT_ID:role/lambda-basic-rol --environment "Variables={STR_MONGO_CONN=$STR_MONGO_CONN, AUTH0_AUDIENCE=$AUTH0_AUDIENCE, AUTH0_DOMAIN=$AUTH0_DOMAIN}"
rm updateQuote.zip updateQuote
