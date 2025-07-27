# Composite

Composite is a structural design pattern that lets you compose objects into tree structures and then work with these 
structures as if they were individual objects.

### Problem

The Composite Pattern solves the problem of how to treat individual objects and groups of objects uniformly,
especially in tree-like structures.

### Components

1. Component (interface/abstract class): Declares common operations. 
2. Leaf: Represents individual objects (no children). 
3. Composite: Represents a group — can have children (leaves or other composites).