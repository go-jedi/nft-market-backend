CREATE OR REPLACE FUNCTION admin_add_balance(_tid BIGINT, _prc NUMERIC)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
	_ub users_base;
BEGIN
	SELECT *
	FROM users
	WHERE tele_id = _tid
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь не найден';
	END IF;

	UPDATE users_base SET balance = balance + _prc WHERE tele_id = _tid;
END;
$function$
