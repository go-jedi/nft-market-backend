CREATE OR REPLACE FUNCTION user_update_premium(_tid BIGINT)
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

	SELECT * FROM
	users_base
	WHERE tele_id = _tid
	INTO _ub;

	IF _ub.is_premium = TRUE THEN
		UPDATE users_base SET is_premium = FALSE WHERE tele_id = _tid;
	ELSE
		UPDATE users_base SET is_premium = TRUE WHERE tele_id = _tid;
	END IF;
END;
$function$
