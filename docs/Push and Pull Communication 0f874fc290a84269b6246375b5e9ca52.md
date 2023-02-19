# Push and Pull Communication

### Push Communication

Server pushes new data to the clients **without them explicitly asking** for it as soon as the data arrive

e.g. Push Notification, **SSE**, WebSocket, pub-sub, WebHooks

### Pull Communication

Client asks server for new data periodically after certain **interval**

This is popularly known as **polling** 

e.g. HTTP Polling Request

**Example**:

Clients asking for updated price of **RELIANCE** stock from NSE

[Pull Example](Push%20and%20Pull%20Communication%200f874fc290a84269b6246375b5e9ca52/Pull%20Example%200485730bc21a453ea9b44608653f7bb3.md)

[Push Example](Push%20and%20Pull%20Communication%200f874fc290a84269b6246375b5e9ca52/Push%20Example%207859c7463f8841fab149fe573fc46584.md)

### Pros and Cons

Pull Communication

- Pros
    1. Easier to implement as no special arrangements are required besides periodic invocation of API.
    2. No persistent connections are required so less memory overhead on server as they don’t require to remember connections for too long.
    3. Easily scalable horizontally as we have stateless connection (for application layer at-least)
- Cons
    1. Delayed updates
    2. Too many un-necessary requests on the server as most of them won’t get any data. This also puts load on various datasources
    3. Cognitive load to calculate the right interval between periodic invokes

Push Communication

- Pros
    1. No-Delay in updates, they are pushed as soon as they are available
    2. Less number of requests on server, as we maintain a persistent connection between client and server
    3. Gives near realtime experience 
- Cons
    1. Complex to scale horizontally for some use-cases (e.g. Group Chat). Vertical scaling won’t work as we have limit for open connections on system level and scaling vertically is more expensive and tedious process.
    2. It’s memory intensive as every connection would take some space in system memory (It can be optimised though)
    3. Requires additional implementations on server side to broadcast updates among plethora of clients.
    

### Long Polling is the midway

- It can be used as an alternative to SSE or WebSockets
- We periodically make request to the server if server have the data (updates) it will immediately send the data and will close the connection immediately.
- If doesn’t have any updates server keeps the connection open (persists) for a certain timeout duration. If in the waiting period it got something it will send it to the client and will also close the connection.
- After getting the connection closed, Client makes a new request immediately

### Use Cases

- Push
    - You want near realtime communication
- Pull
    - You want easier implementation, No fancy implementation for server side
    - You want easier horizontally scalable solution