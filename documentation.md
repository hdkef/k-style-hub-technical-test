# member

### GET /members

query params : limit (int) & page (int)

responses :

    [
        {   
            "id_member":integer,
            "username":string,
            "gender":ENUM('Female','Male')
            "skintype":ENUM('Normal','Dry','Oily')
            "skincolor":ENUM('Type 1','Type 2','Type 3','Type 4','Type 5','Type 6')
        }
    ]

### POST /members

requestBody :

{
    "username":string,
    "gender":ENUM('Female','Male')
    "skintype":ENUM('Normal','Dry','Oily')
    "skincolor":ENUM('Type 1','Type 2','Type 3','Type 4','Type 5','Type 6')
}

### PUT /members/{id} (perlu bearer token)

requestBody :

{
    "username":string,
    "gender":ENUM('Female','Male')
    "skintype":ENUM('Normal','Dry','Oily')
    "skincolor":ENUM('Type 1','Type 2','Type 3','Type 4','Type 5','Type 6')
}

### DELETE /members/{id} (perlu bearer token)

# product

### GET /products/{id}

responses:

    {
        "id_product":integer,
        "name_product":string,
        "price":float,
        "review":[
            {
                "id_review":integer,
                "id_product":integer,
                "id_member":integer,
                "desc":string,
                "member":{   
                    "id_member":integer,
                    "username":string,
                    "gender":ENUM('Female','Male')
                    "skintype":ENUM('Normal','Dry','Oily')
                    "skincolor":ENUM('Type 1','Type 2','Type 3','Type 4','Type 5','Type 6')
                }
            }
        ]
    }

# review

### POST /reviews/{id}/like (perlu bearer token)


## DELETE /review/{id}/like (perlu bearer token)