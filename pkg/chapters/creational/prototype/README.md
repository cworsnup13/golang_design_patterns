## Singleton Pattern

### Intent
Specify the kinds of objects to create using a protoypical instance, and crate new object by copying this prototype. 

### Applicability
- When a system should be independent of how its products are created composed and represented **AND**
    - The classes to instantiate are specified at run-time
    - to avoid building a class hierarchy of factories that parallels the class hierarchy of products
    - when instances of a clas can have one of only a few different combinations of state.

### Participants
- Prototype: Declares an interface for cloning itself.
- ConcreteProtype: immplements an operation for cloning itself
- Client: creates a new object by asking a prototype to clone itself. 


### Golang Notes
- The most difficult part of the prototype approach in go for me was copying interface types. Since they are pointers and there are many
concrete implementations of the interface, type assertion in order to make a copy is complicated.  After review of options I decided to try out
[Gob Encoding](https://blog.golang.org/gobs-of-data) instead of reflection.
- This requires us to register points to all concrete implementations of MapSite with the gob package.