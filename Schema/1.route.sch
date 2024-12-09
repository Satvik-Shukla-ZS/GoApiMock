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

GET /cloud-account/
   CLOUD-ACCOUNT[] :
       id
       content
       created_at
       user
       isPublished
       revenue
       readers
       providers


GET /cloud-account/<>/infra/<>/
   CLOUD-ACCOUNT :
       id

GET /cloud-provider
   CLOUD-PROVIDER[] :
       id
       content
       created_at