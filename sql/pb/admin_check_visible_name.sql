CREATE OR REPLACE FUNCTION admin_check_visible_name(_tid BIGINT)
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
		RAISE EXCEPTION 'администратор не найден';
	END IF;

	IF _a.is_show_name = TRUE THEN
		RETURN TRUE;
	ELSE
		RETURN FALSE;
	END IF;
END;
$function$