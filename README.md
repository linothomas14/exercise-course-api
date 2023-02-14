# ABSENSI API

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

## Routes

| HTTP METHOD                 | POST |           GET            |         PUT          |        DELETE        |
| --------------------------- | :--: | :----------------------: | :------------------: | :------------------: |
| /users                      |  -   |       List of User       |          -           |          -           |
| /users/`<int:id>`           |  -   |      Detail of User      |      UpdateUser      |      DeleteUser      |
| /user-courses               |  -   |    List of UserCourse    |          -           |          -           |
| /user-courses/`<int:id>`    |  -   |   Detail of UserCourse   |   UpdateUserCourse   |   DeleteUserCourse   |
| /courses                    |  -   |      List of Course      |          -           |          -           |
| /courses/`<int:id>`         |  -   |     Detail of Course     |     UpdateCourse     |     DeleteCourse     |
| /course-category            |  -   |  List of CourseCategory  |          -           |          -           |
| /course-category/`<int:id>` |  -   | Detail of CourseCategory | UpdateCourseCategory | DeleteCourseCategory |
| /admins                     |  -   |      List of Admin       |          -           |          -           |
| /admins/`<int:id>`          |  -   |     Detail of Admin      |     UpdateAdmin      |     DeleteAdmin      |

## API Documentation

### List of Endpoints

- [User](#user)
  - [Get AllUser](#get-user)
  - [Get UserDetail](#get-user)
  - [Add User](#add-user)
  - [Delete User](#delete-user)
- [UserCourse](#user-course)
  - [Get AllUserCourse](#get-user-course)
  - [Get UserCourseDetail](#get-user-course)
  - [Add UserCourse](#add-user-course)
  - [Delete UserCourse](#delete-user-course)
- [Course](#course)
  - [Get AllCourse](#get-course)
  - [Get CourseDetail](#get-course)
  - [Add Course](#add-course)
  - [Delete Course](#delete-course)
- [CourseCategory](#course-category)
  - [Get AllCourseCategory](#get-course-category)
  - [Get CourseCategoryDetail](#get-course-category)
  - [Add CourseCategory](#add-course-category)
  - [Delete CourseCategory](#delete-course-category)
- [Admin](#admin)
  - [Get AllAdmin](#get-admin)
  - [Get AdminDetail](#get-admin)
  - [Add Admin](#add-admin)
  - [Delete Admin](#delete-admin)

## User

### Get All User

- Method : GET
- URL : `/users`
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

### Get Mahasiswa by ID

- Method : GET
- URL : `/api/mahasiswa/<int:id>`
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

## Presensi

### Add Presensi

- Method : POST
- URL : `/api/presensi`
- Request body :

```json
{
  "npm": "56419764",
  "minggu": 1,
  "matkul": "Bisnis Informatika"
}
```

- Response body :

```json
{
  "message": "OK",
  "data": {
    "id": 1,
    "npm": "56419764",
    "matkul": "Bisnis Informatika",
    "minggu": 1
  }
}
```

### Get Presensi

- Method : GET
- URL : `/api/presensi?matkul=string&minggu=string`
- Request body : -

- Response body :

```json
{
  "message": "OK",
  "data": {
    "matkul": "Bisnis Informatika",
    "minggu": 1,
    "mahasiswa": [
      {
        "nama": "YULYANO THOMAS DJAYA",
        "npm": "56419764"
      }
    ]
  }
}
```

### Delete Presensi

- Method : DELETE
- URL : `/api/presensi`
- Request body :

```json
{
  "npm": "56419764",
  "minggu": 1,
  "matkul": "Bisnis Informatika"
}
```

- Response body:

```json
{
  "message": "Data berhasil dihapus",
  "data": ""
}
```

## Contributor

- Dwi Pertiwi Ani - 51419933
- Muhammad Arief Rubbyansyah - 54419032
- Shafiah Qonita - 56419004
- Yulyano Thomas Djaya - 56419764
