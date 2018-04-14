## Decorator Pattern

### Intent
Convert the interface of a class into another interface clients expect. Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.  

### Applicability
- You want to use an existing class, and its interface does not match the one you need
- You want to create a reusable class the cooperates with unrelated or unforeseen classes.


### Participants
- Target: defines the domain specific interface that Client uses
- Client: collaborates with objects conforming to the Target interface
- Adaptee: defines an existing interface that needs adapting
- Adapter: adapts the interface of the adaptee to the target interface 