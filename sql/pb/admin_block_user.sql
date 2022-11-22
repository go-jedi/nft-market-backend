CREATE OR REPLACE FUNCTION admin_block_user(_tid BIGINT)
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

	UPDATE users SET is_block = TRUE WHERE tele_id = _tid;
END;
$function$
