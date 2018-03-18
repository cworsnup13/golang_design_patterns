## Composites

#### Motivation

We want to treat objects that are either singular or composed of many components the same. 
For example, drawing a circle or a venn diagram (composed of two circles) both need to be **drawn**.
The client shouldn't need to know the details, just that there exists a draw function. 

This abstraction has the benefit of letting us group things into larger and larger components.

####Golang Implementation Notes
- Golang doesn't have inheritance. I worked around this by having new types have a Composite as the core of its struct.
All operations interacted with the owned composite. This is a bit verbose since we still have to implement each function
to be interface compliant.
- Functions intended for the composite (e.g. Add, Remove) were required at leaf level to implement interface.
- Add / Remove by Pointer: This was an interesting feature of Golang I found. function parameters that expect an interface
 type can receive a pointer type of an implementing type. This would normally fail for regular types but works with interfaces.
 Further research led me to [this article](https://medium.com/@agileseeker/go-interfaces-pointers-4d1d98d5c9c6) which digs 
 into that a bit more. 