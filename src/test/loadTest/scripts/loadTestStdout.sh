#/bin/sh

# for reference and quick test, outputs to stdout

echo "============================================================"
echo "Starting linear load test..."

echo "Load testing sign api:"

CONCURRENCY=4000
HOST="10.0.53.7"
DATA="hello"
DIGEST="LPJNul+wow4m6DsqxbninhsWHlwfp0JecwQzYpOLmCQ="

ddosify -m POST -t http://$HOST:8000/api/sign/k1 -b "$DATA" -n $CONCURRENCY -l linear

echo "Load testing pure-sign api"

ddosify -m POST -t http://$HOST:8000/api/pure-sign/k1 -b "$DIGEST" -n $CONCURRENCY -l linear

echo "============================================================"

echo "Starting incremental load test..."

echo "Load testing sign api"

ddosify -m POST -t http://$HOST:8000/api/sign/k1 -b "$DATA" -n $CONCURRENCY -l incremental

echo "Load testing pure-sign api"

ddosify -m POST -t http://$HOST:8000/api/pure-sign/k1 -b "$DIGEST" -n $CONCURRENCY -l incremental

echo "============================================================"

echo "Starting waved load test..."

echo "Load testing sign api"

ddosify -m POST -t http://$HOST:8000/api/sign/k1 -b "$DATA" -n $CONCURRENCY -l waved

echo "Load testing pure-sign api"

ddosify -m POST -t http://$HOST:8000/api/pure-sign/k1 -b "$DIGEST}" -n $CONCURRENCY -l waved