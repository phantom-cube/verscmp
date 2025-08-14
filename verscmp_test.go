package main

import (
    "testing"
)

func TestVersionCompare(t *testing.T) {
    tests := []struct {
        v1, v2   string
        expected int8
    }{
        {"1.0", "1.0.0", RESULT_LT},
        {"1.0a", "1.0b", RESULT_LT},
        {"1.2", "1.10", RESULT_LT},
        {"v1.2", "v1.10", RESULT_LT},
        {"535.98", "535.146.02", RESULT_LT},
        {"1.2.3", "1.2.3.4", RESULT_LT},
        {"2.0", "2.0", RESULT_EQ},
        {"1.10.0", "1.2.0", RESULT_GT},
        {"a1.2", "a1.10", RESULT_LT},
        {"100", "99", RESULT_GT},
        {"alpha", "beta", RESULT_LT},
    }

    for _, test := range tests {
        result, err := VersionCompare(test.v1, test.v2)
        if err != nil {
            t.Errorf("VersionCompare(%q, %q) returned error: %v", test.v1, test.v2, err)
        }
        if result != test.expected {
            t.Errorf("VersionCompare(%q, %q) = %d, want %d", test.v1, test.v2, result, test.expected)
        }
    }
}

func TestInvalidInput(t *testing.T) {
    _, err := VersionCompare(".1", "1.0")
    if err == nil {
        t.Error("Expected error for version starting with '.'")
    }

    _, err = VersionCompare("1.0.", "1.0")
    if err == nil {
        t.Error("Expected error for version ending with '.'")
    }

    _, err = VersionCompare("1.0#abc", "1.0")
    if err == nil {
        t.Error("Expected error for version with invalid characters")
    }
}