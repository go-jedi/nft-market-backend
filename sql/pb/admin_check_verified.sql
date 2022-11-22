CREATE OR REPLACE FUNCTION admin_check_verified(_tid BIGINT)
	RETURNS BOOLEAN
	LANGUAGE plpgsql
AS $function$
DECLARE
	_ub users_base;
BEGIN
	SELECT *
	FROM users_base
	WHERE tele_id = _tid
	INTO _ub;

	IF _ub.verification = TRUE THEN
		RETURN TRUE;
	ELSE
		RETURN FALSE;
	END IF;
END;
$function$
