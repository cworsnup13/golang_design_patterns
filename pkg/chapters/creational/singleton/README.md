## Singleton Pattern

### Intent
Ensure a class only has one instance and provide a global point of access to it. 

### Applicability
- There must be exactly one instance of a class and accessible to clients
- The sole instance is extensible by subclassing and clients can use an extended instance without modifying code.

### Participants
- Singleton: Defines an instance operation for clients to access.

### Golang notes
- We need to consider goroutines in this implementation. We do want exactly one singleton instance with thread safety.
- Thinking and researching I found a good [blog post](http://marcio.io/2015/07/singleton-pattern-in-go/) on this subject: 