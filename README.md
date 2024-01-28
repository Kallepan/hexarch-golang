# Hexagonal Architecture

## Description

Hexagonal architecture, also known as Ports and Adapters architecture or the Onion architecture, is a software design pattern that promotes a clean separation between the core logic of an application and its external dependencies. It emphasizes making the core business logic of an application independent of external frameworks, libraries, and interfaces, thus increasing its flexibility, testability, and maintainability.

The name "hexagonal" comes from the shape of the architecture when depicted graphically, with the core application logic represented at the center and the external dependencies represented as "ports" on the perimeter.

## Principes

1. Core logic: The core business logic of the application resides at the center of the architecture. This logic is isolated from external concerns and remains independent of specific technologies or implementation details.

2. Ports and adapters: External dependencies such as databases, web frameworks, UI components, and external services are abstracted away through ports and adapters. Ports define the interfaces through which the core logic interacts with these external dependencies, while adapters implement these interfaces and handle the communication with the external systems.

3. Inversion of control: The control flow in hexagonal architecture is inverted compared to traditional architectures. Instead of the core logic depending on external frameworks or libraries, external components depend on interfaces defined by the core logic. This inversion of control decouples the core logic from its dependencies and allows for easier testing and swapping of implementations.

4. Testability: By decoupling the core logic from its external dependencies, hexagonal architecture facilitates unit testing of the core application logic in isolation from external systems. Mocks or stubs can be used to simulate the behavior of external dependencies during testing.

5. Flexibility and maintainability: Hexagonal architecture makes it easier to evolve and maintain software systems over time. Changes to external dependencies, such as switching databases or updating third-party libraries, can be implemented with minimal impact on the core application logic.
