POST :
    id : UUID { minChar=4 , maxChar=20 }
    content : content { minChar=10 , maxChar=20 }
    created_at : datetime { * / * / 2024 }
    user : USER.id {}
    isPublished : bool {}
    revenue : float  { minChar=4 , maxChar=20 }
    readers : int  { minChar=4 , maxChar=20 }