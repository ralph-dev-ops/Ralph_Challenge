#!/usr/bin/env bash

echo "Verifying HTTP redirect"
RESPONSE_CODE=$(curl -o -I -L -s -w "%{http_code}" http://ralph-secnet.us-east-1.elasticbeanstalk.com)
echo "Response Code: $RESPONSE_CODE"

if [ "$RESPONSE_CODE" == "301" ]; then
	echo -e "Redirected!\n"
else
	echo -e "Not Redirected!\n"
fi

echo "Verifying HTML Content"
EXPECTED=`cat infrastructure/index.html`
CONTENT=$(curl -k https://ralph-secnet.us-east-1.elasticbeanstalk.com)

echo -e "Retrieved Content: $CONTENT"

if [ "$EXPECTED" == "$CONTENT" ]; then
	echo "Content Match!"
else
	echo "Content Incorrect!"
fi
