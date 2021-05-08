GOOS=linux go build addNote.go
zip addNote.zip addNote
aws lambda create-function --function-name addNote --zip-file fileb://addNote.zip --handler addNote --runtime go1.x --role arn:aws:iam::$ACCOUNT_ID:role/lambda-basic-rol --environment "Variables={STR_MONGO_CONN=$STR_MONGO_CONN, AUTH0_AUDIENCE=$AUTH0_AUDIENCE, AUTH0_DOMAIN=$AUTH0_DOMAIN}"
rm addNote.zip addNote
