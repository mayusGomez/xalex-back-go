GOOS=linux go build getQuotes.go
zip getQuotes.zip getQuotes
aws lambda create-function --function-name getQuotes --zip-file fileb://getQuotes.zip --handler getQuotes --runtime go1.x --role arn:aws:iam::$ACCOUNT_ID:role/lambda-basic-rol --environment "Variables={STR_MONGO_CONN=$STR_MONGO_CONN, AUTH0_AUDIENCE=$AUTH0_AUDIENCE, AUTH0_DOMAIN=$AUTH0_DOMAIN}"
rm getQuotes.zip getQuotes
