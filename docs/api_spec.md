# API Endpoints


## 1. Create user (Register)
- Endpoint: /api/v1/auth/
- Method: POST
- Headers:
  - Content Type : application/json
- Request
  - Parameter: -
  - Query: -
  - Body:
    ```
    {
        username, -> required
        role, -> required
    }
    ```
- Response:
  - http status: Created/Bad Request/StatusUnprocessableEntity
  - body : 
    ```
    {
      user : {
        username,
        password,
        role,
      }
    }
    ```

## 2. login user
- Endpoint: /api/v1/auth/login
- Method: POST
- Headers:
  - Content Type : application/json
- Request
  - Parameter: -
  - Query: -
  - Body:
    ```
    {
        username -> required
        password -> required
    }
    ```
- Response:
  - http status: Ok/Bad Request/Not Found/Forbidden
  - body : 
    ```
    {
      user : {
        username,
        id,
        token,
      }
    }
    ```

## 3. Validate Token
- Endpoint: /api/v1/auth/:token
- Method: GET
- Headers: -
- Request
  - Parameter: token from login
  - Query: -
  - Body: - 
- Response:
  -  http status: Ok/Unauthorized
  - body : 
    ```
    {
        validate : {
            username,
            is_valid,
            expired,
        }
    }
    ```
