aws apigatewayv2 create-route \
    --api-id 1f4ref09x5 \
    --route-key 'PUT /billing/v1/quotes'
    --target arn:aws:lambda:us-east-1:$ACCOUNT_ID:function:getQuote