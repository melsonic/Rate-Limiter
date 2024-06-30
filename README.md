# Rate-Limiter

### Application Information

#### Token Bucket Algorithm
- Open a terminal(t1) & run either
    - `make run`
    - `make buildnrun`
- To test the scenerio open another terminal(t2) & run
    - `make test_token_bucket`
- In t1, the output shows a line for each request in the following format
    - {Client ID} {Token Count} {Capacity} {Request Allowed/Disallowed}
- Outputs are placed in the `result` folder
