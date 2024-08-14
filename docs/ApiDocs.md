# API Docs

## API lists
1. [Get Posts](#get-posts)
2. [Create Post](#create-post)
3. [Edit Post](#edit-post)
4. [Delete Post](#delete-post)
5. [Sign Up](#sign-up)
6. [Log In](#log-in)

## Get Posts
Retrieve all posts from the database.
```http
GET api/post
```
#### Response
```json
{
    "posts": [
        {
            "id": 21,
            "content": "TEST",
            "createdAt": "2024-08-14T21:22:32.710613+07:00",
            "updatedAt": "2024-08-14T21:22:32.710613+07:00",
            "user": {
                "id": 2,
                "username": "test",
                "displayName": "Test Account"
            }
        }
    ]
}
```

<hr>

## Create Post
Create a new post and require authentication token.
```http
POST api/post/create
```
#### Request Body
```json
{
    "content": "This is my simple social media which allows user to post text, learning space for Next + Go."
}
```

#### Response
```json
{
    "message": "Post created",
    "post": {
        "id": 22,
        "content": "This is my simple social media which allows user to post text, learning space for Next + Go.",
        "createdAt": "2024-08-14T21:41:59.9975086+07:00",
        "updatedAt": "2024-08-14T21:41:59.9975086+07:00",
        "user": {
            "id": 1,
            "username": "Test",
            "displayName": "Admin Account 2"
        }
    }
}
```
<hr>

## Edit Post
Update a post by id and require authentication token.
```http
PUT api/post/{id}
```

#### Request Body
```json
{
    "content": "This is my simple social media which allows user to post text, learning space for Next + Go."
}
```

#### Response
```json
{
    "message": "Post updated",
    "post": {
        "id": 22,
        "content": "[EDIT] This is my simple social media which allows user to post text, learning space for Next + Go.",
        "createdAt": "2024-08-14T21:41:59.997508+07:00",
        "updatedAt": "2024-08-14T21:45:16.9678638+07:00",
        "user": {
            "id": 1,
            "username": "Test",
            "displayName": "Admin Account 2"
        }
    }
}
```
<hr>

## Delete Post
Delete a post by id and require authentication token.
```http
DELETE api/post/{id}
```

#### Response
```json
{
    "message": "Post deleted"
}
```
<hr>

## Sign Up
Create a new user account.
```http
POST api/auth/signup
```

#### Request Body
```json
{
    "username": "test",
    "password": "password",
    "displayName": "Test Account"
}
```

#### Response
```json
{
    "createdAt": "2024-08-14T21:56:17.4564259+07:00",
    "message": "User created",
    "user": {
        "id": 7,
        "username": "Test",
        "displayName": "Admin Account for Learning"
    }
}
```

<hr>

## Log In
Authenticate user and return a token.
```http
POST api/auth/login
```

#### Request Body
```json
{
    "username": "test",
    "password": "password"
}
```

#### Response
```json
{
    "message": "Login success",
    "token": "TOKEN"
}
```

<hr>
