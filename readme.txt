


Endpoints: 

POST: JOBS
    accepts JSON WITH TEST SPECIFICS
    [{
        method: "GET | ",
        url: string
        mode: ENUM("COMPARE")
        equals: string
    }]
GET: JOBS/{ID}

GET: TELEMETRIE/{JOBS_ID}



The jobs accept a testCases object that has different Mode
- COMPARE // compares request response with equals object


