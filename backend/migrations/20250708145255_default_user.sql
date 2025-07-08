DO
$$
    DECLARE
        user_name     TEXT := '{{.default_user}}';
        user_password TEXT := '{{.default_password}}';
    BEGIN
        INSERT INTO users (name, password_hash, created_at, updated_at)
        VALUES (user_name, user_password, NOW(), NOW());
    END
$$;
