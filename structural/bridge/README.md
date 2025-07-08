# Bridge

The Bridge design pattern is a structural design pattern that aims to decouple an abstraction from its implementation so
that the two can vary independently.

### Problem 
* When you have an entity that can be classified along two or more independent dimensions (e.g., a Shape can be 
Circle/Square and Red/Blue; a Device can be TV/Radio and BasicRemote/AdvancedRemote). Without Bridge, you might create 
classes like RedCircle, BlueCircle, RedSquare, BlueSquare, etc., leading to M * N classes. 
* Adding a new dimension of variation (e.g., a new Green color or a new DVDPlayer device) can require modifying many 
existing classes.

### Components 

1. Abstraction (Interface in Go, often with an embedding struct):
   * Defines the abstraction's client-facing interface. Maintains a reference to an object of the Implementor interface
    type. It delegates most of its work to the Implementor object. Example: Printer interface. 
2. Refined Abstraction(s) (Concrete Structs implementing Abstraction):
   * Extends the Abstraction's interface, providing more specific control or functionality. Example: Mac, 
   Windows. 
3. Implementor (Interface in Go):
   * Defines the interface for the implementation classes. This interface provides only primitive operations that the 
   Abstraction needs. It does not need to mirror the Abstraction's interface directly. Example: Printer interface with 
   methods like PrintFile().
4. Concrete Implementor(s) (Concrete Structs implementing Implementor):
   * Implement the Implementor interface. These are the actual underlying objects that perform the specific operations. 
   Example: Epson, Hp.