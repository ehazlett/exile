#!/bin/bash
set -e
ADDR=127.0.0.1:8080
export TEST_URL=http://127.0.0.1:8080

execute() {
    >&2 echo "++ $@"
        eval "$@"
}

./exile server -c exile.conf.sample -l $ADDR &
EXILE_PID=$!
sleep 2

echo "$EXILE_PID"

TESTS="api cli"

for TEST in $TESTS; do
    execute bats test/integration/$TEST.bats || true
done

kill $EXILE_PID
