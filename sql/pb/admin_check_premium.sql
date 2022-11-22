CREATE OR REPLACE FUNCTION admin_check_premium(_tid BIGINT)
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

	IF _ub.is_premium = TRUE THEN
		RETURN TRUE;
	ELSE
		RETURN FALSE;
	END IF;
END;
$function$
