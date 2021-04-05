


Endpoints: 

POST: JOBS
    accepts JSON WITH TEST SPECIFICS
    [{
        method: "GET | ",
        url: string
        mode: ENUM("COMPARE", "SNAPSHOT")
        equals: string
    }]
GET: JOBS/{ID}

GET: TELEMETRIE/{JOBS_ID}



The jobs accept a testCases object that has different Mode
- COMPARE // compares request response with equals object
    - equals // represents object to compare in json
- ruleset
    - finite set of rules to apply to testCase

