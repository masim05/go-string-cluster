# go-string-cluster
A simple utility to cluster strings into groups of "similar" ones.

# Problem
Sometimes you need to analyze error messages and cluster them by the root
cause. The error messages are not always exactly the same, since they might
contain dynamic data. One can use Levenshtein distance to address it, however
the messages might have different lengths, so it can be not easy to find a
proper threshold.

# Solution
This command line tool uses `weighted` Levenshtein distance:
$D_{LW} = \dfrac{D_{L}(s_1, s_2)}{\|s_1\| + \|s_2\|}$

and clusters strings according to it.

# Usage
Installation:
```bash
go install github.com/masim05/go-string-cluster/scluster@latest
```
Usage:
```bash
% echo "123\n1234\n1123\nqwert" | scluster
Line clusters:
(3)	: 123
(1)	: qwert
```
Output shows two cluster, the first one contains 3 lines, the second one - only 1.

# TODO
- [ ] implement `ClusterStringChannel` which will accept channel of strings and
 pefrom clusterisation without saving all the input in memory
- [ ] support two modes of the CLI tool: regular one (low memory consumption
 using channels) and verbose one, which shows content of each cluster (high
 memory consumption using `ClusterStringSet`)
