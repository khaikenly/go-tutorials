1
=
```json
[
    ["User_001", 0, 1000],
    ["User_002", 500, 2000]
]
```

2
=
```json
[
    ["User_001", 0, 1000],
    ["User_002", 500, 2000],
    ["User_003", 2500, 3000],
    ["User_004", 400, 1400]
]
```

3
=
```json
[
    ["User_001", 0, 1000],
    ["User_002", 500, 2000],
    ["User_003", 2500, 3000],
    ["User_004", 400, 1400],
    ["User_001", 1100, 1800],
    ["User_005", 1200, 1400],
    ["User_006", 500, 2400],
    ["User_003", 2100, 2300]
    
]
```

4
=
```json
[
    ["User_001", 0, 1000],
    ["User_002", 500, 2000],
    ["User_003", 2500, 3000],
    ["User_001", 1100, 1800],
    ["User_002", 600, 2400],
    ["User_002", 700, 1800]
]
```



Concurrent Learners
===================
A popular type of learning available though Go1 is watching streamed videos. Learners view videos as streams from our CDN (content delivery network). 

For this challenge, we'll be looking at when streams are being watched based on start time and end time. The dataset contains a list of user IDs, start and stop timestamps for each stream watched.

We're interested in knowing what was the maximum number of concurrent streams within the dataset. A stream is concurrent with another when the start and end times over lap.

Return the maximum number of concurrent streams.

Example input:
```json
[
    ["User_001", 0, 1000],
    ["User_002", 500, 2000],
    ["User_003", 2500, 3000],
    ["User_004", 400, 1400]
]
```

Example output:
```json
"The maximum number of concurrent streams is 3."
```