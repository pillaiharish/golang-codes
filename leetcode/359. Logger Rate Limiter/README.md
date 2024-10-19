# 359 Logger Rate Limiter

To solve LeetCode 359 Logger Rate Limiter, we need to design a logger system that limits printing of the same message based on a time threshold.


## Problem:
You need to implement the Logger class:

shouldPrintMessage(timestamp int, message string) bool: Returns true if the message should be printed (not printed in the last 10 seconds); otherwise, return false.


## Approach:
We use a hash map to store the last printed timestamp for each message and check if 10 seconds have passed since the last print.


## Explanation:
logs is a map where keys are the messages, and values are the last timestamp they were printed.
In ShouldPrintMessage, we check if the message was printed in the last 10 seconds using the timestamp. If it was, return false; otherwise, update the timestamp and return true.


## Time Complexity:
O(1) for both shouldPrintMessage operations (since map operations are O(1)).


## Space Complexity:
O(n), where n is the number of unique messages stored in the map.