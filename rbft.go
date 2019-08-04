package main

// This file will hold the main logic for the randomized consensus protocol

import (
    "github.com/cnnrznn/channel"
)

func sum(a []int) int {
    sum := 0

    for _, e := range a {
        sum += e
    }

    return sum
}

func faulty(ch channel.Channel) int {
    f := len(ch.Peers)
    return (f-1) / 3
}

func round(ch channel.Channel, v int, r int) (int, bool) {
    message_count := make([]int, 2)
    echo_count := make([][2]int, len(ch.Peers))
    decision := false
    var msg channel.Msg

    for sum(message_count) < (2 * faulty(ch) + 1) {
        msg = ch.Recv()

        message_count[0]++
    }

    if message_count[0] > message_count[1] {
        v = 0
    } else {
        v =1
    }

    for _, e := range message_count {
        if e > 2 * faulty(ch) {
            decision = true
        }
    }

    return v, decision
}

