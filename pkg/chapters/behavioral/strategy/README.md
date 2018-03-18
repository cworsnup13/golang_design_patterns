## Strategy Pattern

### Intent
Define a family of algorithms, encapsulate each one and make them interchangeable. 
Strategy lets the algorithm vary independently from clients that use it.

### Applicability
- Many related classes differ only in their behavior
- You need different variants of an algorithm
- an algorithm uses data the client shouldn't know about
- instead of many conditionals, move related conditional branches into their own strategy class

### Participants
- Strategy: Interface common to all supported algorithms
- ConcreteStrategy: implements the strategy interface
- Context: forwards requests from clients to its strategy

