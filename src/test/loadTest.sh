#/bin/sh

echo "Starting linear load test..."

echo "Load testing sign api"

ddosify -m POST -t http://localhost:8000/api/sign/k1 -b "{\"data\": \"hello\"}" -n 10000 -l linear

echo "Load testing pure-sign api"

ddosify -m POST -t http://localhost:8000/api/pure-sign/k1 -b "{\"digest\": \"LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=\"}" -n 10000 -l linear


echo "Starting incremental load test..."

echo "Load testing sign api"

ddosify -m POST -t http://localhost:8000/api/sign/k1 -b "{\"data\": \"hello\"}" -n 10000 -l incremental

echo "Load testing pure-sign api"

ddosify -m POST -t http://localhost:8000/api/pure-sign/k1 -b "{\"digest\": \"LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=\"}" -n 10000 -l incremental


echo "Starting waved load test..."

echo "Load testing sign api"

ddosify -m POST -t http://localhost:8000/api/sign/k1 -b "{\"data\": \"hello\"}" -n 10000 -l waved

echo "Load testing pure-sign api"

ddosify -m POST -t http://localhost:8000/api/pure-sign/k1 -b "{\"digest\": \"LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ=\"}" -n 10000 -l waved