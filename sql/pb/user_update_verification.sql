CREATE OR REPLACE FUNCTION user_update_verification(_tid BIGINT)
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

	SELECT *
	FROM users_base
	WHERE tele_id = _tid
	INTO _ub;

	UPDATE users_base SET verification = TRUE WHERE tele_id = _tid;
END;
$function$
