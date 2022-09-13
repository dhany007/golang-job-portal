# Job Portal

## Description
Make API for Job Portal. All development use `git flow`.

## Prerequisite
  - Golang
  - Postgres
  - Git Flow
  -

## Architecture
![img-architecture](architecture-golang.png)
source: https://medium.com/easyread/golang-clean-archithecture-efd6d7c43047

## Scopes
| Ticket ID | Ticket Title | User Story |
|---|---|---|
| [JP-01](readme.md#user-register) | User register | As a User, i can register as company or candidate |
| [JP-02](readme.md#user-login) | User login, refresh access token and logout | As a User, i can login as company or candidate, refresh access token and logout |
| [JP-03](readme.md#update-profil-company) | Update profil company | As a company, i can update my profil |
| [JP-04](readme.md#update-profil-candidate) | Update profil candidate | As a candidate, i can update my profil |
| [JP-05](readme.md#get-list-company) | Get list company | As a candidate, i can see list of company |
| [JP-06](readme.md#get-detail-company) | Get detail company | As a candidate, i can see detail of company |
| [JP-07](readme.md#review-company) | Review company | As a candidate, i can review company with rating |
| [JP-08](readme.md#get-list-review-company) | Get list review company | As a candidate, i can get list of all reviews of company |
| [JP-09](readme.md#get-list-dresscode-company) | Get list company dresscode, benefits, and size | As a company, i can get list code dresscode, benefits and size |

## Api Specs

### User Register
  - Desctription : User for register
  - Method : `POST`
  - Endpoint : `/users/register`
  - Parameter : -
  - Request
    ```
    {
      "email": string,required,email,
      "password": string,required,length(6|32),
      "role": int,required,range(1|2)
    }
    ```
  - Response
    - Success
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        },
        "data": {
          "id": string,uuid,
          "email": string,
          "is_active": int,
          "role": int,
          "created_at": string,time,utc,
          "modified_at": string,time,utc
        }
      }
      ```
    - Failed
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        }
      }
      ```

### User Login
  - Description: User for login, user will get access-token for access all endpoint and set to cookie
  - Method : `POST`
  - Endpoint : `/users/login`
  - Parameter : -
  - Request
    ```
    {
      "email": string,required,email,
      "password": string,required,length(6|32),
    }
    ```
  - Response
    - Success
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        },
        "data": {
          "refresh-token": string,uuid,
          "access-token": string,uuid
        }
      }
      ```
    - Failed
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        }
      }
      ```

### User Refresh Token
  - Description : Endpoint for refresh access-token from refresh-token. If refresh-token expired, user should login again.
  - Method : `POST`
  - Endpoint : `/users/refresh-token`
  - Parameter : -
  - Request: -
  - Response
    - Success
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        },
        "data": {
          "access-token": string,uuid
        }
      }
      ```
    - Failed
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        }
      }
      ```

### Update Profil Company
  - Description : Endpoint for update company, check if company_id is same with id user login
  - Method : `POST`
  - Endpoint : `/companies/:companyID`
  - Parameter :
  - Request: -
    ```
    {
      "email": string,
      "name": string,required,
      "description": string,required,
      "address": string,required,
      "website": string,
      "phone_number": string,required,
      "telp_number": string,
      "profil_picture_url": string,
      "dress": int,
      "size": int,
      "benefit": [
        {"id": int},
      ]
    }
    ```
  - Response
    - Success
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        },
        "data": {
          "email": "string",
          "name": "string",
          "description": "string",
          "address": "string",
          "website": "string",
          "phone_number": "string",
          "telp_number": "string",
          "profil_picture_url": "string",
          "dress": "string",
          "size": "strings",
          "benefit": [
            {
              "id": int,
              "value": "string"
            }
          ],
          "review": {
            "rating": float,
            "count_review": int
          },
          "created_at": "string",
          "modified_at": "string"
        }
      }
      ```
    - Failed
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        }
      }
      ```

### Update Profil Candidate

### Get List Company
  - Description : Endpoint for get list companies sort by rating ascending.
  - Method : `GET`
  - Endpoint : `/companies`
  - Parameter :
    - `page`: optional, default 1
    - `item_per_page`: optional, default 10
  - Request: -
  - Response
    - Success
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        },
        "data": {
          "page": int,
          "total_page": int,
          "item_per_page": int,
          "total_data": int,
          "companies": [
            {
              "id": "string",
              "name": "string",
              "rating": int,
              "count_review": int
            }
          ]
        }
      }
      ```
    - Failed
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        }
      }
      ```

### Get Detail Company
  - Description : Endpoint for detail company
  - Method : `GET`
  - Endpoint : `/companies/detail/:companyID`
  - Parameter : -
  - Request: -
  - Response
    - Success
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        },
        "data": {
          "email": "string",
          "name": "string",
          "description": "string",
          "address": "string",
          "website": "string",
          "phone_number": "string",
          "telp_number": "string",
          "profil_picture_url": "string",
          "dress": "string",
          "size": "strings",
          "benefit": [
            {
              "id": int,
              "value": "string"
            }
          ],
          "review": {
            "rating": float,
            "count_review": int
          },
          "created_at": "string",
          "modified_at": "string"
        }
      }
      ```
    - Failed
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        }
      }
      ```

### Review Company
  - Description : Endpoint candidate to give review company
  - Method : `POST`
  - Endpoint : `/companies/reviews`
  - Parameter : -
  - Request:
    ```
    {
      "company_id": "string,required",
      "candidate_id": "string,required",
      "rating": int,numeric,range(1|3),
      "review":"string,required"
    }
    ```
  - Response
    - Success
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        },
        "data": {
          "id": int,
          "company_id": "string",
          "candidate_id": "string",
          "rating": int,
          "review": "string"
        }
      }
      ```
    - Failed
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        }
      }
      ```

### Get List Review Company
  - Description : Endpoint for get list revies company based on id and sorting by new added.
  - Method : `GET`
  - Endpoint : `/companies/reviews/{companyID}`
  - Parameter :
    - `page`: optional, default 1
    - `item_per_page`: optional, default 10
  - Request: -
  - Response
    - Success
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        },
        "data": {
          "page": int,
          "total_page": int,
          "item_per_page": int,
          "total_data": int,
          "reviews": [
            {
              "id": int,
              "candidate_id": "string",
              "rating": int,
              "review": "string"
            }
          ]
        }
      }
      ```
    - Failed
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        }
      }
      ```


### Get List Dresscode Company
  - Description : Endpoint for get list code dresscode company.
  - Method : `GET`
  - Endpoint : `/companies/dress-codes`
  - Parameter : -
  - Request: -
  - Response
    - Success
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        },
        "data": [
          {
            "id": int,
            "value": "string"
          }
        ]
      }
      ```
    - Failed
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        }
      }
      ```

### Get List Benefits Company
  - Description : Endpoint for get list code benefits company.
  - Method : `GET`
  - Endpoint : `/companies/benefit-codes`
  - Parameter : -
  - Request: -
  - Response
    - Success
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        },
        "data": [
          {
            "id": int,
            "value": "string"
          }
        ]
      }
      ```
    - Failed
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        }
      }
      ```

### Get List Size Company
  - Description : Endpoint for get list code size company.
  - Method : `GET`
  - Endpoint : `/companies/size-codes`
  - Parameter : -
  - Request: -
  - Response
    - Success
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        },
        "data": [
          {
            "id": int,
            "value": "string"
          }
        ]
      }
      ```
    - Failed
      ```
      {
        "status": int,
        "message": {
          "en": "string",
          "id": "string"
        }
      }
      ```
