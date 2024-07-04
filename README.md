### What is a Rate-Limiter 🤔

Rate Limiter is a Decision engine in web applications that decides whether a client request should be fulfilled or rejected. A rate limiter tracks the request rate against each client and if that exceeds a certain threshold, it just rejects the upcoming requests from that particular client.

### Advantages of using a Rate Limiter 🔥 
- **Securing Applications**
    - It helps servers to survive brute force attacks. When an web application is under attack, it faces huge number of requests, hence in that situation if a rate limiter is available, it can control the request rate from such clients.

---

### Rate Limiting Algorithms ⚒️
There are different algorithms for this purpose, each having its own advantages. Here are the four algorithms that are implemented in this application.

### Token Bucket Algorithm 
##### Important Definitions
1. **Token** : A `Token` is like a lifeline for a client request.
2. **Token_Rate** : Rate at which new tokens are added. 
3. **Threshold** : Maximum value of `Token`.

##### Algorithm Points
- Against each client, a `Token` count is stored. It is like a lifeline for that client. When a request arrives, if the `Token` count is greater than 0, it allows the request, otherwise it discards the request.
- Tokens are added to the `Token` variable at a rate `Token_Rate`.
- When a token are added, if the `Token` count becomes greater than the `Threshold` value, it discards the `Token`.

### Fixed Window Counter
##### Algorithm Points
- The timeline is divided into small time chunks of equal size, called a `Window`.
- For each client, it stores the number of requests the client made in a particular time `Window`.
- If the request count exceeds a `Threshold` value, it discards the request.
- This process is repeated for each window.

### Sliding Window Log
##### Algorithm Steps
- It is a bit dynamic in nature compared to Fixed Window Counter.
- It also has the concept of `Window` of a particular duration(let's say `Window_Size`). 
- Against each client, it stores the timestamps at which the client made a request.
- When a request arrives, it discards all the previous timestamps that lies outside of the `Window` i.e. (now - `Window_Size` ➡️  now)
- Then it counts the number of available timestamps in the current `Window`, if it exceeds a particular `Threshold` value, then it discards the request.

### Sliding Window Counter
##### Algorithm Points
- Sliding Window Counter is a combination of Fixed Window Counter and Sliding Window Log.
- It divides the timeline into `Window` of same sizes.
- For each client, it stores the request count for both the Current window and the Previous window.
- To calculate the estimated request count for the current `Sliding Window`, it it's 37% through in the Current Window, it takes the 63% of the requests made in the Previous Window to make up for it.
- Then it compares the value against a particular Threshold value and takes a decision whether to allow or discard the request.
---

### Application Information

- Open a terminal(t1) & run either
  - `make run`
  - `make buildnrun`
- On running the above command, you will be prompted with algorithm options.
- Select the algorithm (1, 2, 3 ...)
- 👉 To **run the tests**, you need to install [k6](https://k6.io/docs/get-started/installation/)

#### Token Bucket Algorithm

- To test the scenerio open another terminal(t2) & run
  - `make test_token_bucket`
- In t1, the output shows a line for each request in the following format
  - {Client ID} {Token Count} {Capacity} {Request Allowed/Disallowed}
- Outputs are placed in the `result` folder

#### Fixed Window Algorithm

- To test the fixed window algorithm, open another terminal(t2) & run
  - `make test_fixed_window`
- In t1, the output shows a line for each request in the following format
  - {Client ID} {CurrentWindowRequestCount} {Request Allowed/Disallowed}
- Outputs are placed in the `result` folder

#### Sliding Window Log Algorithm

- To test the sliding window log algorithm, open another terminal(t2) & run
  - `make test_sliding_window_log`
- In t1, the output shows a line for each request in the following format
  - {Client ID} {CurrentWindowTimeStampCount} {TimeStampCountThreshold} {Request
    Allowed/Disallowed}
- Outputs are placed in the `result` folder

#### Sliding Window Counter Algorithm

- To test the sliding window log algorithm, open another terminal(t2) & run
  - `make test_sliding_window_counter`
- In t1, the output shows a line for each request in the following format
  - {Client ID} {EstimatedCurrentWindowCount} {CurrentWindowThreshold} {Request
    Allowed/Disallowed}
- Outputs are placed in the `result` folder
