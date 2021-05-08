GOOS=linux go build addNote.go
zip addNote.zip addNote
aws lambda update-function-code --function-name addNote --zip-file fileb://addNote.zip
rm addNote.zip addNote