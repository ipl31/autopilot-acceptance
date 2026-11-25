# Interval membership

`window.Contains(first, last, value)` reports membership in a closed interval of
Go `int` values. Both endpoints belong to the interval, equal bounds describe a
singleton, and reversed bounds describe an empty interval. The contract applies
across the supported platform's full `int` range.
