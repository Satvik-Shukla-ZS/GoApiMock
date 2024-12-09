APPLICATION :
    id : UUID { minChar=4 , maxChar=20 }
    content : content { minChar=10 , maxChar=20 }
    created_at : datetime { * / * / 2024 }

CLOUD-PROVIDER :
    id : UUID { minChar=4 , maxChar=20 }
    content : content { minChar=10 , maxChar=20 }
    created_at : datetime { * / * / 2024 }
    application : APPLICATION { minChar=2 , maxChar=5 }

CLOUD-ACCOUNT :
    id : UUID { minChar=4 , maxChar=20 }
    content : content { minChar=10 , maxChar=20 }
    created_at : datetime { * / * / 2024 }
    user : USER.id {}
    isPublished : bool {}
    revenue : float  { minChar=4 , maxChar=20 }
    readers : int  { minChar=4 , maxChar=20 }
    providers : CLOUD-PROVIDER { minChar=2 , maxChar=5 }


ORGANIZATION :
    id : UUID { minChar=4 , maxChar=20 }
    content : content { minChar=10 , maxChar=20 }
    created_at : datetime { * / * / 2024 }
    application : APPLICATION { minChar=2 , maxChar=5 }