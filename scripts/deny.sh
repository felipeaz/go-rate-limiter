#!/bin/bash
target=${1:-localhost:8080/test}
while true # loop forever, until ctrl+c pressed.
do
	for i in $(seq 5) # perfrom the inner command 100 times.
	do
		curl $target > /dev/null & # send out a curl request, the & indicates not to wait for the response.
	done

	wait # after the requests are sent out, wait for their processes to finish before the next iteration.
done