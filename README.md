# simpleserver
simple http server

## Task
You should create your mongoDB service in the localhost, fulfill user story below and demo it.


### User story
As a client, it is able to get last 10 entities from API `/test/all`

### Acceptance Criteria
- Test code: test these scenario 
  - There are no items in the `test` collection
  - There are more than 10 items in the `test` collection
  - There are less than 10 items in the `test` collection
- Use query parameter like this `/test/all?last=10`
- Error handle: API should response error with match status code, the error response format is:
  ```json
  {
    "err":"reason"
  }
  ```