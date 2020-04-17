# Strmetric
Package `strmetric` provides functions to measure string similarity and distance for Go lang.

[![Build Status](https://travis-ci.com/enginyoyen/strmetric.svg?branch=master)](https://travis-ci.com/enginyoyen/strmetric)
## Metrics 
Currently, package provides following similarity metrics;

* [Dice / Sorensen](http://en.wikipedia.org/wiki/Dice%27s_coefficient)
* [Hamming](http://en.wikipedia.org/wiki/Hamming_distance)
* [Jaro](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance)
* [Jaro-Winkler](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance)
* [Levenshtein](http://en.wikipedia.org/wiki/Levenshtein_distance)

## Examples
#### Dice / Sorensen Metric
Dice-Sørensen algorithm is used to gauge the similarity of two samples.
```
strmetric.DiceSorensonMetric("night", "nacht") // 0.6
```

#### Hamming Metric
The Hamming distance measures the minimum number of substitutions required to change one string into the other, or the minimum number of errors that.
Used for error detection.

```
strmetric.HammingMetric("11011001", "10011101") // 2
```


#### Jaro Metric
Jaro Similarity is the edit distance between two strings. The higher the Jaro distance for two strings is, the more similar the strings are.

```
strmetric.JaroMetric("JELLYFISH", "SMELLYFISH") // 0.896296
```
#### Jaro-Winkler Metric
Uses the Jaro similarity but takes both string prefix into an account(up-to 4 characters) and factors into a score.

```
strmetric.JaroMetric("martha", "marhta") // 0.944444
strmetric.JaroWinklerMetric("martha", "marhta") // 0.961111
```
#### Levenshtein
Levenshtein distance is a string metric for measuring the difference between two sequences.
Used for spell checkers, error detection/correction in optical character recognition, etc.
```
strmetric.LevenshteinMetric("kitten", "sitting") // 3
```

# Licence 
Use of this source code is governed by an MIT license that can be found in the LICENSE file.
