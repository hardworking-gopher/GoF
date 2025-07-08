# GoF design patterns in Golang

A practical collection of classic Gang of Four (GoF) Design Patterns implemented in idiomatic Go (Golang). This repository serves as a personal quick reference to recall how these patterns look in Go.

## What are GoF design patterns?

Design Patterns are common, reusable solutions to recurring problems in software design. The "Gang of Four" (GoF) refers to Erich Gamma, Richard Helm, Ralph Johnson, and John Vlissides, who authored the seminal book "Design Patterns: Elements of Reusable Object-Oriented Software." Their work codified 23 classic patterns, providing a shared vocabulary and framework for developers to discuss and apply proven solutions.

## How are they grouped?

The GoF patterns are categorized into three main types based on their purpose:

1.  **Creational Patterns:** These patterns deal with **object creation mechanisms**, trying to create objects in a manner suitable to the situation while increasing flexibility and reuse of code. They abstract the instantiation process, making the system independent of how its objects are created, composed, and represented.

2.  **Structural Patterns:** These patterns deal with **class and object composition**. They describe how objects and classes can be combined to form larger structures, promoting flexibility and efficiency. Structural patterns focus on organizing different classes and objects to form structures that are larger than the individual classes, allowing them to collaborate more effectively.

3.  **Behavioral Patterns:** These patterns deal with **algorithms and the assignment of responsibilities between objects**. They describe the communication patterns between objects, focusing on how objects interact with each other to perform a task. Behavioral patterns aim to ensure objects can easily communicate and carry out tasks, increasing flexibility in communication.