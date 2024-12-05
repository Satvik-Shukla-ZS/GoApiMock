GET /

GET /user/<>/details
   USER :
       name
       id

GET /user
   USER[] :
       name
       id
       email

GET /posts
   POST[] :
       id
       content
       created_at
       user
       isPublished
       revenue
       readers

GET /posts/<>
   POST :
       id
       content
       created_at
       user
       isPublished
       revenue
       readers