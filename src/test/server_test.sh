SCRIPT_PATH=`dirname $0`
HSM_SERVICE=`find ${SCRIPT_PATH}/.. -name 'hsm-service' -type f -executable`
KEY_PATH=${SCRIPT_PATH}/testrsaprivkey.pem
KEY2_PATH=${SCRIPT_PATH}/testrsaprivkey1.pem

rm nohup.out
nohup ${HSM_SERVICE} serve --k1 ${KEY_PATH} --k2 ${KEY2_PATH}  &
PID=$!
echo "Server is running with pid ${PID}"

# wait for server to come up, loop till health check is 200
while true; do
    sleep 1
    STATUS=`curl -s -o /dev/null -w "%{http_code}" http://localhost:8000/api/health`
    if [ "${STATUS}" == "200" ]; then
        break
    fi
done


echo -n "Sign API Test"
RESULT=`curl -X POST http://localhost:8000/api/sign/k1 -s -d "{\"data\": \"${DATA}\"}" | yq .result`
if [ ! $? == 0 ]; then
    echo " [FAILED]"
    exit 1
fi
echo " [PASS]"

echo -n "Pure Sign API Test"
DATA=`echo -n "hello" | openssl dgst -sha256 -binary - | base64 -w 0`
RESULT=`curl -X POST http://localhost:8000/api/pure-sign/k1 -s -d "{\"digest\": \"${DATA}\"}" | yq .result`
if [ ! $? == 0 ]; then
    echo " [FAILED]"
    exit 1
fi
echo " [PASS]"

echo -n "Bad Request Test"
RESULT=`curl -X POST http://localhost:8000/api/pure-sign/k1 -s -d "hello" | yq .result`
echo " [PASS]"
sleep 0.5
kill $PID