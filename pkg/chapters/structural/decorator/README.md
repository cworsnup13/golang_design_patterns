## Decorator Pattern

### Intent
Attach additional responsibilities to an object dynamically. Decorators provide a flexible alternative to subclassing 
for extending functionality. 

### Applicability
- To add responsibilities to individual objects dynamically and transpartenly.
- For responsibilities that can be withdrawn

### Participants
- Component: defines te interface for objects that can have responsibilities added dynamically
- ConcreteComponent: implements the component interface
- Decorator: maintains a reference to a component object and defines an interface that conforms to Component's interface
- ConcreteDecorator: adds responsibilities to the component

