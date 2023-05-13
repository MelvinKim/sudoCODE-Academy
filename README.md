# sudoCODE-Academy
This is a system  where students can enroll to different courses. Built using the Microservices architecture and run on Kubernetes

### Services:
 1. Users
 2. Payments
 3. Courses
 4. Notifications
 5. each request has a request ID that will help to track the flow of request across the different services

### When a creating a new student account:
1. check for duplicate account
2. if there's no duplicate account, we charge the user
3. then create a user profile and assign a course(s)
4. send the student notifications via email, for confirmation (include the courses that they have been assigned)
5. for phase-1, everything works synchronously, with the hope that nothing fails in between :)

#### Users service
1. creates a user account
2. forwards the user to the payment service

#### Payments service
1. charges the user an annual subscription fee
2. forwards the user to the courses service

#### Courses service
1. creates a user profile
2. assigns the user the courses they had selected when signing up

#### Notifications service
1. sends notification to the user via email: welcoming them to sudoCODE academy and listing the selected courses

### API requirements
#### User
- GET /api/v1/users
- GET /api/v1/users/123
- POST /api/v1/users
- DELETE /api/v1/users/123

#### Payment
- GET /api/v1/payments
- GET /api/v1/payments/123
- POST /api/v1/payments
- DELETE /api/v1/payments/123

#### Course
- GET /api/v1/courses
- GET /api/v1/courses/123
- POST /api/v1/courses
- DELETE /api/v1/courses/123

#### Notification
- GET /api/v1/notifications
- GET /api/v1/notifications/123
- POST /api/v1/notifications
- DELETE /api/v1/notifications/123
