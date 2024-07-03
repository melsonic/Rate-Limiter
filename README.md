# Rate-Limiter

### Application Information
- Open a terminal(t1) & run either
    - `make run`
    - `make buildnrun`
- On running the above command, you will be prompted with algorithm options.
- Select the algorithm (1, 2, 3 ...)

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
    - {Client ID} {CurrentWindowTimeStampCount} {TimeStampCountThreshold} {Request Allowed/Disallowed}
- Outputs are placed in the `result` folder
