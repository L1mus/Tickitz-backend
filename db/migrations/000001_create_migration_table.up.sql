DO
$$
    BEGIN
        CREATE TYPE role_user AS ENUM ('user', 'admin');
    EXCEPTION
        WHEN duplicate_object THEN NULL;
    END
$$;

DO
$$
    BEGIN
        CREATE TYPE condition_ticket AS ENUM
            (
                'active',
                'not_active'
                );
    EXCEPTION
        WHEN duplicate_object THEN NULL;
    END
$$;

DO
$$
    BEGIN
        CREATE TYPE condition_paid AS ENUM ('paid', 'not_paid');
    EXCEPTION
        WHEN duplicate_object THEN NULL;
    END
$$;

DO
$$
    BEGIN
        CREATE TYPE status_transaction AS ENUM ('completed', 'pending', 'failed');
    EXCEPTION
        WHEN duplicate_object THEN NULL;
    END
$$;

DO
$$
    BEGIN
        CREATE TYPE type_seat AS ENUM ('regular', 'love_nest');
    EXCEPTION
        WHEN duplicate_object THEN NULL;
    END
$$;
