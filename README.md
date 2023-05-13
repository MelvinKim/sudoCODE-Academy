# sudoCODE-Academy
This is a system  where students can enroll to different courses. Built using the Microservices architecture and run on Kubernetes

### Services:
 1. Users
 2. Payments
 3. Courses
 4. Notifications

### When a creating a new student account:
1. check for duplicate account
2. if there's no duplicate account, we charge the user
3. then create a user profile and assign a course(s)
4. send the student notifications via email, for confirmation (include the courses that they have been assigned)
5. for phase-1, everything works synchronously, with the hope that nothing fails in between :)
