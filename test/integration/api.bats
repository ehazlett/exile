#!/usr/bin/env bats

load helpers

@test "api: index" {
    run curl -s $TEST_URL/
    [ "$status" -eq 0  ]
}

@test "api: sign" {
    CERT=/tmp/exile-test.pem
    if [ -e $CERT ]; then rm $CERT; fi
    run curl -s -d @test/certs/csr.json $TEST_URL/sign -o $CERT
    [ "$status" -eq 0  ]
    run openssl x509 -in $CERT -noout -text
    echo "$output" | grep "Issuer: O=exile"
    [ "$status" -eq 0  ]
    run openssl x509 -in $CERT -noout -text
    echo "$output" | grep "DNS:node-00"
    [ "$status" -eq 0  ]
    rm $CERT
}
