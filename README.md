# Go-Masterclass
Includes all exercises for mastering and understand golang in depth

### **1. Basic Syntax and Constructs**
- [x]  **FizzBuzz in Go:** Write a program to print numbers from 1 to 100, replacing multiples of 3 with "Fizz", multiples of 5 with "Buzz", and multiples of both with "FizzBuzz".
- [x]  **Prime Numbers Generator:** Generate all prime numbers up to `N` using different approaches (e.g., Sieve of Eratosthenes, basic iteration).
- [x]  **Reverse a String:** Implement a function to reverse a string, considering UTF-8 characters.

---

### **2. Concurrency and Goroutines**

- [ ]  **Ping-Pong Using Channels:** Implement a program where two goroutines send "ping" and "pong" messages back and forth using channels.
- [ ]  **Worker Pool:** Create a worker pool to process a list of tasks concurrently, ensuring all tasks are completed.
- [ ]  **Rate Limiter:** Design and implement a rate limiter using Go channels.

---

### **3. Interfaces and Structs**

- [ ]  **Shape Interface:** Define an interface for shapes with methods like `Area()` and `Perimeter()`, and implement it for different shapes like circle, rectangle, and triangle.
- [ ]  **Dynamic Polymorphism:** Implement a system where you can add new types dynamically to satisfy an interface.
- [ ]  **Custom Error Handling:** Create a custom error type that includes additional information, like a timestamp or error severity.

---

### **4. Advanced Data Structures**

- [ ]  **Linked List Implementation:** Implement a singly and doubly linked list in Go.
- [ ]  **Binary Search Tree:** Build a binary search tree with methods for insertion, deletion, and traversal (in-order, pre-order, post-order).
- [ ]  **LRU Cache:** Implement an LRU (Least Recently Used) cache with concurrency control.

---

### **5. File Handling**

- [ ]  **Log Parser:** Write a program to read and parse a log file, extracting specific fields or filtering entries based on conditions.
- [ ]  **File Watcher:** Implement a file watcher that monitors a directory for changes and prints updates in real time.
- [ ]  **Large File Reader:** Efficiently read a file that is too large to fit into memory.

---

### **6. Networking**

- [ ]  **HTTP Server:** Write a basic HTTP server using the `net/http` package with endpoints for CRUD operations.
- [ ]  **Chat Application:** Build a simple TCP-based chat server where multiple clients can communicate with each other.
- [ ]  **Web Scraper:** Build a web scraper that fetches and parses data from a website.

---

### **7. Testing and Benchmarking**

- [ ]  **Custom Testing Framework:** Create a simple testing framework similar to Go's `testing` package.
- [ ]  **Benchmark Sorting Algorithms:** Compare the performance of different sorting algorithms using Go's benchmarking tools.

---

### **8. Reflection and Generics**

- [ ]  **Dynamic Struct Validation:** Use reflection to validate struct fields based on custom tags.
- [ ]  **Generic Stack:** Implement a generic stack using Go generics for push, pop, and peek operations.

---

### **9. Memory Management and Optimization**

- [ ]  **Custom Memory Pool:** Create a custom memory pool to allocate and reuse objects efficiently.
- [ ]  **Detecting Race Conditions:** Write a program with a potential race condition and resolve it using synchronization techniques like `sync.Mutex`.

---

### **10. Real-World Applications**

- [ ]  **Task Scheduler:** Build a task scheduler that executes tasks at specific times using Go's `time` package.
- [ ]  **Simple Key-Value Store:** Create a basic in-memory key-value store with persistence to disk.
- [ ]  **JWT Authentication:** Implement a system to generate and validate JWTs for authentication