CREATE OR REPLACE FUNCTION admin_update_min_price(_tid BIGINT, _minp NUMERIC)
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

	UPDATE admins SET minim_price = _minp WHERE tele_id = _tid;
END;
$function$
