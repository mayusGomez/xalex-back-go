GOOS=linux go build getQuote.go
zip getQuote.zip getQuote
aws lambda create-function --function-name getQuote --zip-file fileb://getQuote.zip --handler getQuote --runtime go1.x --role arn:aws:iam::$ACCOUNT_ID:role/lambda-basic-rol --environment "Variables={STR_MONGO_CONN=$STR_MONGO_CONN, AUTH0_AUDIENCE=$AUTH0_AUDIENCE, AUTH0_DOMAIN=$AUTH0_DOMAIN}"
rm getQuote.zip getQuote
