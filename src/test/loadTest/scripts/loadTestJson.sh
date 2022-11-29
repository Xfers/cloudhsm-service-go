#/bin/sh

# data from the command line
CONCURRENCY=$1
ENDPOINT=$2
DATA=$3 #it can be digest

ddosify -m POST -t $ENDPOINT -b "$DATA" -n $CONCURRENCY -o stdout-json
