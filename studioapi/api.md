### Base REST API

- POST ***/v1/users/signin*** - log in for an existing user
- GET ***/v1/lessons?date=:date&time=:time*** - list of lessons
- POST ***/v1/lessons/:lessonID*** - book a specific lesson
- POST ***/v1/orders/:lessonID*** - pay for a specific lesson

### Product-first approach
By the product-first approach I would add some another base scenarios: like 
- POST ***/v1/users/signup*** - create a new user,

possibility of change date/time, canceling order etc.

### Tests
I would write unit test for API by mocking server-side using packages like
https://golang.org/pkg/net/http/httptest

A also believe that integration tests for API are required, I would do it by writing separate scenarios for each endpoint. 
