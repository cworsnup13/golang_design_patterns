## Abstract Factory Pattern

### Intent
Provide an ingerface for creating families of related or dependent objects without specifying their concrete classes. 

### Applicability
- A system should be independent of how its products are created, composed and represented.
- A system should be configured with one of multiple families of products.
- A family of related product objects is designed to be used together and you need to enforce this constraint
- You want to provide a class library of products and you want to reveal just their interfaces not their implementations.

### Participants
- Abstract Factory: Declares an interface for operations that create abstract product objects
- Concrete Factory: implemtns the operations to create conrete product objects
- Abstract Product: Declares an interface for a type of product object
- Concrete Product: Defines a product object ot be created by the corresponding concrete factory. Implements the abstract product interface
- Client: uses only interfaces declared by Abstract Factory and Abstract Product.

### Golang notes
- We need to consider goroutines in this implementation. We do want exactly one singleton instance with thread safety.
- Thinking and researching I found a good [blog post](http://marcio.io/2015/07/singleton-pattern-in-go/) on this subject: 