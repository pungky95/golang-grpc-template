INSERT INTO
    "public"."users" (
        "id",
        "created_at",
        "updated_at",
        "deleted_at",
        "name",
        "email",
        "phone",
        "password",
        "status"
    )
VALUES
    (
        gen_random_uuid(),
        NOW(),
        NOW(),
        NULL,
        'Pungky Kurniawan Kristianto',
        'pungkykurniawankr@gmail.com',
        '62812235626668',
        '$argon2id$v=19$m=65536,t=3,p=4$OdmYUi817aQnwlJ7X5x3wQ$JxdwygEUz5kYBAZckN90N/MIiAo4VGuchL9G3mw+Idk',
        'ACTIVE'
    );