<?php

// I check for:
// Readability
// Understandability
// Performance
// Error handling
// Bugs & Logical flaws
// DRY-Principle -> Don't write redundant code
// SOP-Principle -> Each thing does its own thing
// KISS-Principle -> Try to keep it simple
// SOLID-Principle -> https://en.wikipedia.org/wiki/SOLID
// POLA-Principle -> Code should work as one would expect
// YAGNI-Principle -> Dont code stuff you might not need

echo fibonacci(10) . PHP_EOL;

try {
    echo factorial(5) . PHP_EOL;
} catch (Exception $e) {
    echo $e;
}

echo isPrime(43) . PHP_EOL;

printHello("Alice");

// Removed abbreviations to make code easier to read and understand
// Added return type
// Code wouldn't work if length was value 1. Added a check to catch that case.
function fibonacci($number) : string {
    if ($number <= 0) {
        return 0;
    }

    if ($number == 1) {
        return 1;
    }

    $fibonacci = [0, 1];
    for ($index = 2; $index < $number; $index++) {
        $fibonacci[] = $fibonacci[$index - 1] + $fibonacci[$index - 2];
    }

    return implode(' ', $fibonacci);
}

// Removed abbreviations to make code easier to read and understand
// for-loop was flawed which lead to an incorrect factorial. Fixed it by replacing = with <=
// Added error handling for negative numbers
// Added return type
function factorial($number) : int {
    if ($number < 0) {
        throw new Exception("'invalid input: can't calculate factorial with negative numbers'");
    }

    if ($number == 0) {
        return 1;
    }

    $result = 1;
    for ($index = 1; $index <= $number; $index++) {
        $result *= $index;
    }

    return $result;
}

// Renamed function according to best practices https://www.php.net/manual/en/userlandnaming.rules.php
// Removed abbreviations to make code easier to read and understand
// Added return type
function isPrime($number) : bool {
    if ($number < 2) {
        return false;
    }

    for ($index = 2; $index < $number / 2; $index++) {
        if ($number % $index == 0) {
            return false;
        }
    }

    return true;
}

// Renamed function according to best practices https://www.php.net/manual/en/userlandnaming.rules.php
// Instead of implementing error handling, an empty string will instead produce a hello message without personalization
// Used sprintf instead of + string concatenation, although sprintf may be less performant.
function printHello($name): void {
    if (strlen($name) > 0) {
        echo sprintf("Hello, %s!", $name);
    } else {
        echo "Hello!";
    }
}
