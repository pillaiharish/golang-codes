# Parallel Square Calculator

This Go project demonstrates **fan-out/fan-in concurrency**:
- **Fan-out:** Multiple workers read from the same jobs channel to parallelize computation.
- **Fan-in:** Results from all workers are merged into a single results channel for further processing.

---

## 🚀 Features
- Uses **goroutines** to run workers in parallel.
- Demonstrates **sync.WaitGroup** for goroutine synchronization.
- Uses **channel closing patterns** to gracefully stop workers.
- Calculates the **square of the first 20 numbers**.

---

## 💡 How It Works

### **Worker Function**
Each worker:
- Reads jobs from a `jobs` channel.
- Computes the square.
- Sends the result into a `results` channel.
- Signals completion via `wg.Done()`.

### **Main Function Steps**
1. Create `jobs` and `results` channels.
2. Launch `N` workers.
3. Send jobs to workers.
4. Wait for workers to finish.
5. Collect and print results.

---

## 🔄 Concurrency Pattern Used
The code uses the **Fan-out/Fan-in** pattern:
```
Producer → jobs → [Worker 1]
[Worker 2] → results → Consumer (main)
[Worker 3]
```

---


- **Fan-out:** Multiple workers pull from a single `jobs` channel.
- **Fan-in:** All workers push into the same `results` channel.

---

## 📝 Example Output
```
square: 1
square: 4
square: 9
square: 16
...
square: 400
```

---

## 🧠 Key Takeaways
- Use **`sync.WaitGroup`** to track worker goroutines.
- **Close channels** to signal termination.
- Separate **job feeding**, **worker processing**, and **result collection** stages.

