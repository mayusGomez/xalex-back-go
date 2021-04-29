
aws lambda update-function-code --function-name getServices --zip-file fileb://getServices.zip --environment "Variables={STR_MONGO_CONN=$STR_MONGO_CONN, AUTH0_AUDIENCE=$AUTH0_AUDIENCE, AUTH0_DOMAIN=$AUTH0_DOMAIN}"
