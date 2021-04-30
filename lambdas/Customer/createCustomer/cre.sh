GOOS=linux go build createCustomer.go
zip createCustomer.zip createCustomer
aws lambda create-function --function-name createCustomer --zip-file fileb://createCustomer.zip --handler createCustomer --runtime go1.x --role arn:aws:iam::$ACCOUNT_ID:role/lambda-basic-rol --environment "Variables={STR_MONGO_CONN=$STR_MONGO_CONN, AUTH0_AUDIENCE=$AUTH0_AUDIENCE, AUTH0_DOMAIN=$AUTH0_DOMAIN}"
rm createCustomer.zip createCustomer
