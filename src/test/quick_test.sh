#/bin/sh

SCRIPT_DIR=`dirname $0`

cd $SCRIPT_DIR/..

if [ -f "hsm-service" ]; then
    rm hsm-service
fi

go build -o hsm-service .


echo "GROUND TRUTH SIGN:"
echo -n "hello" | openssl dgst -sign test/testrsaprivkey.pem -sha256 | base64 -w 0
echo "\n"


echo "SERVICE SIGN:"
./hsm-service sign -k test/testrsaprivkey.pem -s "hello" 
echo "\n"

echo "SERVICE PURE SIGN:"
echo -n "hello" | openssl dgst -sha256 -binary | ./hsm-service pure-sign -k test/testrsaprivkey.pem
echo "\n"
