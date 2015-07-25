#!/usr/bin/env bats

load helpers

@test "cli: show info" {
    run exile
    [ "$status" -eq 0  ]
    [[ ${lines[0]} =~ "NAME:"  ]]
}

@test "cli: version" {
    run exile -v
    [ "$status" -eq 0  ]
    [[ ${lines[0]} =~ "version"  ]]
}
