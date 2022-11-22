CREATE OR REPLACE FUNCTION admin_change_min_user(_tid BIGINT, _minp NUMERIC)
	RETURNS void
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

	UPDATE users SET minim_price = _minp WHERE tele_id = _tid;
END;
$function$
