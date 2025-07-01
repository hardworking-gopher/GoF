# Singleton

The Singleton pattern is a creational design pattern that ensures a class has only one instance, while providing a 
global access point to this instance. It's useful when you need exactly one object to coordinate actions across the
system, like a database connection, a logger, or a configuration manager.

# Core Concept
The main goals of the Singleton pattern are to:

* Restrict Instantiation: Guarantee that only one instance of a class can be created.

* Provide Global Access: Offer a single, well-known point of access to that instance.

* This is typically achieved by making the constructor private and providing a static method that handles the instance creation and retrieval.