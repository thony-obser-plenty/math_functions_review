<?php

function fibonacci($n) {
    if ($n <= 0) {
        return [];
    }

    $fib = [0, 1];
    for ($i = 2; $i < $n; $i++) {
        $fib[] = $fib[$i - 1] + $fib[$i - 2];
    }
    return $fib;
}

function factorial($n) {
    if ($n == 0) {
        return 1;
    }

    $result = 1;
    for ($i = 1; $i < $n; $i++) {
        $result *= $i;
    }
    return $result;
}

function is_prime($num) {
    if ($num < 2) {
        return false;
    }

    for ($i = 2; $i < $num / 2; $i++) {
        if ($num % $i == 0) {
            return false;
        }
    }
    return true;
}

function print_hello($name) {
    echo "Hello, " . $name . "!";
}
