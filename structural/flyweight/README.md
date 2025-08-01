# Flyweight

Flyweight is a structural design pattern that lets you fit more objects into the available amount of RAM by sharing 
common parts of the state between multiple objects instead of keeping all the data in each object.

### Problem

This constant data of an object is usually called the intrinsic state. It lives within the object; other objects
can only read it, not change it. The rest of the object’s state, often altered “from the outside” by other objects,
is called the extrinsic state. The Flyweight pattern suggests that you stop storing the extrinsic state inside the 
object. Instead, you should pass this state to specific methods which rely on it. Only the intrinsic state stays within
the object, letting you reuse it in different contexts. As a result, you’d need fewer of these objects since they only
differ in the intrinsic state, which has much fewer variations than the extrinsic.

### Components

1. 