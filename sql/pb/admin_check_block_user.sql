CREATE OR REPLACE FUNCTION admin_check_block_user(_tid BIGINT)
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

	IF _u.is_block = TRUE THEN
		RETURN TRUE;
	ELSE
		RETURN FALSE;
	END IF;
END;
$function$