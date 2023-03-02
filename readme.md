## Environment

memerlukan,
DB_USERNAME --> username database
DB_PASS --> password dari username database tersebut
DB_HOST --> host database
DB_PORT --> port database
DB_NAME --> nama database
JWT_SECRET --> secret dari jwt
APP_PORT --> port dimana app berjalan

untuk JWT token dapat mengakses dummylogin dengan path /auth/login dengan mengirim
request POST dgn payload id_member berupa integer dan total jam expired berupa integer

example:

    {
        "id_member" : 1
        "exp":2
    }

