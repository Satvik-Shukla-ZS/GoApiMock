USER :
    name : string { minChar=4 , maxChar=20 }
    email : email { minChar=5 , maxChar=20 }
    password : string { minChar=6, maxChar=20 }
    role : string { ADMIN , USER , GUEST }
    created_at : datetime { * / * / 2024 }

