# #359 Logger Rate Limiter — 90-second cold recall card

[Full problem and reasoning](README.md) · [Commented code](main.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md)

**Trigger:** Repeated event decisions depend on the **last accepted** time of each independent key.

**Invariant:** `logs[message]` is the timestamp of that message's latest **allowed**, not latest received, call.

| Step | Reconstruct | Why |
|---|---|---|
| 1 | Put `logs map[string]int` in `Logger` | State must outlive individual method calls |
| 2 | `Constructor() Logger` returns a struct with `make(map[string]int)` | A nil map can be read but panics on assignment |
| 3 | Check `previous, exists := this.logs[message]` | Distinguish never seen from a genuine timestamp zero |
| 4 | If exists and `timestamp-previous < 10`, return false | Reject while waiting; do not extend the window |
| 5 | Save `logs[message] = timestamp`, return true | Only an accepted arrival restarts that key's cooldown |

**Trace:** `foo@1` true; `bar@2` true; `foo@3` false; `foo@10` false; `foo@11` true. Independent keys do not block each other.

**Your Go foundations to revisit:** A local map inside the method resets each call; the constructor initializes a persistent instance; `Logger{}` has a nil map; `Constructor()` returns a **value**; pointer receiver updates instance state. Two `Logger` instances must not share one map by accident.

**Complexity:** average O(1) per call, O(number of distinct messages) retained space. Not automatically thread-safe or bounded in a production stream.

**Cold drill:** Implement `Constructor` and `ShouldPrintMessage` with the exact t+9/t+10 boundary; explain why rejected calls never write to the map.

Confidence: `—` · Last cold recall: `—` · Next recall: `—`
