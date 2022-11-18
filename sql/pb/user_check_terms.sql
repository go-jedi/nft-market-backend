CREATE OR REPLACE FUNCTION user_check_terms(_tid BIGINT)
    RETURNS BOOLEAN
    LANGUAGE plpgsql
AS $function$
DECLARE
    _u users;
BEGIN
    SELECT *
    FROM users
    WHERE tele_id = _tid
    INTO _u;

    IF _u.id ISNULL THEN
        RAISE EXCEPTION 'пользователь не найден';
    END IF;
    IF _u.is_terms = TRUE THEN
        RETURN TRUE;
    ELSE
        RETURN FALSE;
    END IF;
END;
$function$