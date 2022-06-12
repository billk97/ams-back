# ams-back

Access management system backend repo responsible for handling the business logic of ACA-PY

## Projects class diagram

Mermaid diagram.

```mermaid
classDiagram
Resource "0..*" -->  "1" Permission : Requires
Resource "0..*" -->  "1" Action : Requires
ResourceType "0..*" -->  "1" Action 
Permission "1..*" --> "0..*" Action
Employee "0..*" -->  "0..*" Permission : Has
class Admin {
    - int id
    - String userName
    - String password
}
class Employee {
    - int id
    - String firstName
    - String lastName
    - String jobTitle
    - Date dirthDate
    - Sex sex
    - String email
    - String mobileNumber
    - String didConnectionId
}
class Permission {
    - int id
    - String title
}
class Resource {
    - int id
    - String alias
    - String description
}
class Action {
    <<enumaration>>
    ACCESS, READ, WRITE, DELETE
}
class Sex {
    <<enumaration>>
    MALE, FEMALE
}
class ResourceType {
    - int id,
    - int label
}
```

### What a credential may look like

```json
{
    "employee": {
        "firstName": "Vasileios",
        "LastName": "Konstantinou",
        "jobTitle": "Software developer",
        "dirthDate": "2022-06-12T13:44:31Z",
        "sex": "male",
        "email": "a@a.gr",
        "mobileNumber": "+306977177481",
        "permisions": [
            {
                "id": 1,
                "alias": "MEETING_ROOM_1",
                "action": ["ACCESS"]
            },
            {
                "id": 2,
                "alias": "CAFETERIA",
                "action": ["ACCESS"]
            },
            {
                "id": 2,
                "alias": "S3",
                "path": "/"
                "action": ["READ", "WRITE", "DELETE"]
            },
        ],
    },

}
```