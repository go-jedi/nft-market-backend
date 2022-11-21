CREATE OR REPLACE FUNCTION user_check_admin(_tid BIGINT)
	RETURNS BOOLEAN
	LANGUAGE plpgsql
AS $function$
DECLARE
	_a admins;
BEGIN
	SELECT *
	FROM admins
	WHERE tele_id = _tid
	INTO _a;

	IF _a.id ISNULL THEN
		RETURN FALSE;
	ELSE
		RETURN TRUE;
	END IF;
END;
$function$