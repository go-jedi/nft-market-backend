CREATE OR REPLACE FUNCTION user_check_auth(_tid BIGINT)
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
        RETURN FALSE;
    ELSE
        RETURN TRUE;
    END IF;
END;
$function$