1
=
```json
{
  "Learner-0001": [
    "Course-0001"
  ],
  "Learner-0002": [
    "Course-0002"
  ],
  "Learner-0003": [
    "Course-0003"
  ],
  "Learner-0004": [
    "Course-0004"
  ]
}
```

2
=
```json
{
  "Learner-0001": [
    "Course-0001",
    "Course-0002",
    "Course-0003"
  ],
  "Learner-0002": [
    "Course-0002",
    "Course-0003",
    "Course-0004"
  ]
}
```

3
=
```json
{
  "Learner-0001": [
    "Course-0001",
    "Course-0002",
    "Course-0003"
  ]
}
```

4
=
```json
{
  "Learner-0001": [
    "Course-0001",
    "Course-0002",
    "Course-0003",
    "Course-0001",
    "Course-0002",
    "Course-0003",
  ],
  "Learner-0002": [
    "Course-0002",
    "Course-0003",
    "Course-0004"
  ]
}
```

5
=
```json
{
  "Learner-0001": [
    "Course-0001",
    "Course-0002",
    "Course-0003"
  ],
  "Learner-0002": [
    "Course-0002",
    "Course-0003",
    "Course-0004"
  ],
  "Learner-0003": [
    "Course-0004",
    "Course-0005",
    "Course-0006"
  ],
  "Learner-0004": [
    "Course-0005",
    "Course-0006",
    "Course-0007"
  ]
}
```

6
=
```json
{
  "Learner-0001": [
    "Course-0001",
    "Course-0002",
    "Course-0003"
  ],
  "Learner-0002": [
    "Course-0002",
    "Course-0003"
  ],
  "Learner-0003": [
    "Course-0002",
    "Course-0003",
    "Course-0004"
  ]
}
```


Learners and Courses
====================

Learners are people using Go1 to learn something. Courses are videos, quizzes, interactive experiences, etc. 

For this challenge, the dataset contains a map of LearnerIds and a list of CourseIds that the Learner's have completed.

We'd like to know about courses that have only been completed by a single learner.

Return a list of CourseId's that have only been completed by 1 Learner.

Example input:
```json
{
  "Learner-0001": [
    "Course-0001",
    "Course-0002",
    "Course-0003"
  ],
  "Learner-0002": [
    "Course-0002",
    "Course-0003",
    "Course-0004"
  ]
}
```

Example output:
```json
[ "Course-0001", "Course-0004" ]
```