# Course Exercise API

API ini ditujukan untuk memehuni Technical Test Backend Developer Intern

## Table of Contents

- [Setup](#setup)
- [Routes](#routes)
- [API Documentation](#api-documentation)
- [Contributor](#contributor)

## Setup

Untuk menjalankan API ini diperlukan bahasa Go dan MySQL sebagai database

- buat file .env, sesuai dengan .env-example
- jalankan perintah `go run main.go`
- creds akun admin :

```json
{
  "name": "Admin",
  "email": "admin@gmail.com",
  "password": "admin123"
}
```

## Routes

| HTTP METHOD                   |       POST        |           GET            |         PUT          |        DELETE        |
| ----------------------------- | :---------------: | :----------------------: | :------------------: | :------------------: |
| /users                        |         -         |       List of User       |      UpdateUser      |          -           |
| /users/`<int:id>`             |         -         |      Detail of User      |          -           |      DeleteUser      |
| /users/profile                |         -         |        GetProfile        |          -           |          -           |
| /user-courses                 |   EnrollCourse    |    List of UserCourse    |          -           |    UnenrollCourse    |
| /user-courses/`<int:id>`      |         -         |   Detail of UserCourse   |          -           |          -           |
| /courses                      |     AddCourse     |      List of Course      |          -           |          -           |
| /courses/`<int:id>`           |         -         |     Detail of Course     |     UpdateCourse     |     DeleteCourse     |
| /course-categories            | AddCategoryCourse |  List of CourseCategory  |          -           |          -           |
| /course-categories/`<int:id>` |         -         | Detail of CourseCategory | UpdateCourseCategory | DeleteCourseCategory |
| /admins                       |   RegisterAdmin   |      List of Admin       |          -           |          -           |
| /admins/`<int:id>`            |         -         |     Detail of Admin      |     UpdateAdmin      |     DeleteAdmin      |
| /login                        |     LoginUser     |            -             |          -           |          -           |
| /login-admin                  |    LoginAdmin     |            -             |          -           |          -           |
| /register                     |   RegisterUser    |            -             |          -           |          -           |

## API Documentation

### List of Endpoints

- [Auth](#auth)
  - [Login](#login)
  - [Login Admin](#login-admin)
  - [Register](#register)
- [User](#user)
  - [Get AllUser](#get-all-user)
  - [Get UserDetail](#get-user-by-id)
  - [Get Profile](#get-user-profile)
  - [Update User](#update-user)
  - [Delete User](#delete-user)
- [UserCourse](#user-course)
  - [Get AllUserCourse](#get-all-user-course)
  - [Get UserCourseDetail](#get-user-course-by-id)
  - [Enroll Course](#enroll-course)
  - [Unenroll Course](#unenroll-course)
- [Course](#course)
  - [Get AllCourse](#get-all-course)
  - [Get CourseDetail](#get-course-by-id)
  - [Add Course](#add-course)
  - [Update Course](#update-course)
  - [Delete Course](#delete-course)
- [CourseCategory](#course-category)
  - [Get AllCourseCategory](#get-all-course-category)
  - [Get CourseCategoryDetail](#get-course-category-by-id)
  - [Add CourseCategory](#add-course-category)
  - [Update CourseCategory](#update-course-category)
  - [Delete CourseCategory](#delete-course-category)
- [Admin](#admin)
  - [Get AllAdmin](#get-all-admin)
  - [Get AdminDetail](#get-admin-by-id)
  - [Add Admin](#add-admin)
  - [Update Admin](#update-admin)
  - [Delete Admin](#delete-admin)

## Auth

### Login

- Method : POST
- URL : `/login`
- Token : -
- Request body :

```json
{
  "email": "string",
  "password": "string"
}
```

- Response body :

```json
{
  "message": "string",
  "data": {
    "token": "string"
  }
}
```

### Login Admin

- Method : POST
- URL : `/login-admin`
- Token : -
- Request body :

```json
{
  "email": "string",
  "password": "string"
}
```

- Response body :

```json
{
  "message": "string",
  "data": {
    "token": "string"
  }
}
```

### Register

- Method : POST
- URL : `/register`
- Token : -
- Request body :

```json
{
  "email": "string",
  "name": "string",
  "password": "string"
}
```

- Response body :

```json
{
  "message": "string",
  "data": {
    "id": "int",
    "email": "string",
    "name": "string"
  }
}
```

## User

### Get All User

- Method : GET
- URL : `/users`
- Token : `adminToken`
- Request body : -
- Response body :

```json
{
  "message": "OK",
  "data": [
    {
      "id": "int",
      "email": "string",
      "name": "string"
    }
  ]
}
```

### Get User Profile

- Method : GET
- URL : `/users/profile`
- Token: `token`
- Request body : -
- Response body :

```json
{
  "message": "OK",
  "data": {
    "id": "int",
    "email": "string",
    "name": "string"
  }
}
```

### Get User by ID

- Method : GET
- URL : `/users/<int:id>`
- Token: `tokenAdmin`
- Request body : -
- Response body :

```json
{
  "message": "OK",
  "data": {
    "id": "int",
    "email": "string",
    "name": "string"
  }
}
```

### Update User

- Method : PUT
- URL : `/users`
- Token : `tokenUser`
- Request body :

```json
{
  "email": "string",
  "name": "string",
  "password": "string"
}
```

- Response body :

```json
{
  "message": "string",
  "data": {
    "id": "int",
    "email": "string",
    "name": "string"
  }
}
```

### Delete User

- Method : DELETE
- URL : `/users/<id:int>`
- Token : `tokenAdmin`
- Request body : -
- Response body :

```json
{
  "message": "User id `id` was deleted",
  "data": {}
}
```

## Course

### Get All Course

- Method : GET
- URL : `/courses`
- Token : `userToken / adminToken`
- Request body : -
- Response body :

```json
{
  "message": "OK",
  "data": [
    {
      "id": "int",
      "title": "string",
      "course_category_id": "int",
      "course_category": {
        "id": "int",
        "name": "string"
      }
    }
  ]
}
```

### Get Course By ID

- Method : GET
- URL : `/courses`
- Token : `userToken / adminToken`
- Request body : -
- Response body :

```json
{
  "message": "OK",
  "data": {
    "id": "int",
    "title": "string",
    "course_category_id": "int",
    "course_category": {
      "id": "int",
      "name": "string"
    }
  }
}
```

### Add Course

- Method : POST
- URL : `/courses/<id:int>`
- Token : `adminToken`
- Request body :

```json
{
  "title": "string",
  "course_category_id": "int"
}
```

- Response body :

```json
{
  "message": "OK",
  "data": {
    "id": "int",
    "title": "string",
    "course_category_id": "int",
    "course_category": {
      "id": "int",
      "name": "string"
    }
  }
}
```

### Update Course

- Method : PUT
- URL : `/courses/<id:int>`
- Token : `adminToken`
- Request body :

```json
{
  "title": "string",
  "course_category_id": "int"
}
```

- Response body :

```json
{
  "message": "OK",
  "data": {
    "id": "int",
    "title": "string",
    "course_category_id": "int",
    "course_category": {
      "id": "int",
      "name": "string"
    }
  }
}
```

### Delete Course

- Method : DELETE
- URL : `/courses/<id:int>`
- Token : `adminToken`
- Request body : -

- Response body :

```json
{
  "message": "Course id `id` was deleted",
  "data": {}
}
```

## Course Category

### Get All Course Category

- Method : GET
- URL : `/course-categories`
- Token : `userToken / adminToken`
- Request body : -
- Response body :

```json
{
  "message": "OK",
  "data": [
    {
      "id": "int",
      "name": "string"
    }
  ]
}
```

### Get Course Category By ID

- Method : GET
- URL : `/course-categories`
- Token : `userToken / adminToken`
- Request body : -
- Response body :

```json
{
  "message": "OK",
  "data": [
    {
      "id": "int",
      "name": "string"
    }
  ]
}
```

### Add Course Category

- Method : POST
- URL : `/course-categories`
- Token : `adminToken`
- Request body :

```json
{
  "name": "string"
}
```

- Response body :

```json
{
  "message": "OK",
  "data": [
    {
      "id": "int",
      "name": "string"
    }
  ]
}
```

### Update Course Category

- Method : PUT
- URL : `/course-categories/<id:int>`
- Token : `adminToken`
- Request body :

```json
{
  "name": "string"
}
```

- Response body :

```json
{
  "message": "string",
  "data": {
    "id": "int",
    "name": "string"
  }
}
```

### Delete Course Category

- Method : DELETE
- URL : `/course-categories/<id:int>`
- Token : `adminToken`
- Request body : -

- Response body :

```json
{
  "message": "Data has been deleted",
  "data": {}
}
```

## User Course

### Get All User Course

- Method : GET
- URL : `/user-courses`
- Token : `adminToken`
- Request body : -
- Response body :

```json
{
  "message": "string",
  "data": {
    "id": "int",
    "user_id": "int",
    "User": {
      "id": "int",
      "email": "string",
      "name": "string"
    },
    "course_id": "int",
    "Course": {
      "id": "int",
      "title": "string",
      "course_category_id": "int",
      "course_category": {
        "id": "int",
        "name": "string"
      }
    }
  }
}
```

### Get User Course By ID

- Method : GET
- URL : `/user-courses`
- Token : `adminToken`
- Request body : -
- Response body :

```json
{
  "message": "string",
  "data": {
    "id": "int",
    "user_id": "int",
    "User": {
      "id": "int",
      "email": "string",
      "name": "string"
    },
    "course_id": "int",
    "Course": {
      "id": "int",
      "title": "string",
      "course_category_id": "int",
      "course_category": {
        "id": "int",
        "name": "string"
      }
    }
  }
}
```

### Enroll Course

- Method : POST
- URL : `/user-courses`
- Token : `token / adminToken`
- Request body :

```json
{
  "user_id": 3,
  "course_id": 8
}
```

- Response body :

```json
{
  "message": "string",
  "data": {
    "id": "int",
    "user_id": "int",
    "User": {
      "id": "int",
      "email": "string",
      "name": "string"
    },
    "course_id": "int",
    "Course": {
      "id": "int",
      "title": "string",
      "course_category_id": "int",
      "course_category": {
        "id": "int",
        "name": "string"
      }
    }
  }
}
```

### UnEnroll Course

- Method : DELETE
- URL : `/user-courses`
- Token : `token / adminToken`
- Request body :

```json
{
  "user_id": 3,
  "course_id": 8
}
```

- Response body :

```json
{
  "message": "string",
  "data": {
    "id": "int",
    "user_id": "int",
    "User": {
      "id": "int",
      "email": "string",
      "name": "string"
    },
    "course_id": "int",
    "Course": {
      "id": "int",
      "title": "string",
      "course_category_id": "int",
      "course_category": {
        "id": "int",
        "name": "string"
      }
    }
  }
}
```

## Admin

### Get All Admin

- Method : GET
- URL : `/admins`
- Token : `adminToken`
- Request body : -
- Response body :

```json
{
  "message": "string",
  "data": [
    {
      "id": "int",
      "email": "string",
      "name": "string"
    }
  ]
}
```

### Get Admin By ID

- Method : GET
- URL : `/admins/<id:int>`
- Token : `adminToken`
- Request body : -
- Response body :

```json
{
  "message": "string",
  "data": {
    "id": "int",
    "email": "string",
    "name": "string"
  }
}
```

### Add Admin

- Method : POST
- URL : `/admins`
- Token : `adminToken`
- Request body :

```json
{
  "name": "string",
  "email": "string",
  "password": "string"
}
```

- Response body :

```json
{
  "message": "string",
  "data": {
    "id": "int",
    "email": "string",
    "name": "string"
  }
}
```

### Update Admin

- Method : PUT
- URL : `/admins/<id:int>`
- Token : `adminToken`
- Request body :

```json
{
  "name": "string",
  "email": "string",
  "password": "string"
}
```

- Response body :

```json
{
  "message": "string",
  "data": {
    "id": "int",
    "email": "string",
    "name": "string"
  }
}
```

### DELETE Admin

- Method : DELETE
- URL : `/admins/<id:int>`
- Token : `adminToken`
- Request body : -

- Response body :

```json
{
  "message": "Admin id `id` was deleted",
  "data": {}
}
```

## Contributor

- Yulyano Thomas Djaya - 56419764
