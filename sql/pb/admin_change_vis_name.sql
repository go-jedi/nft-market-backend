CREATE OR REPLACE FUNCTION admin_change_vis_name(_tid BIGINT)
	RETURNS void
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
		UPDATE admins SET is_show_name = FALSE WHERE tele_id = _tid;
	ELSE
		UPDATE admins SET is_show_name = TRUE WHERE tele_id = _tid;
	END IF;
END;
$function$