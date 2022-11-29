CREATE OR REPLACE FUNCTION admin_withdraw_refuse(_tid BIGINT, _uid character varying)
	RETURNS BOOLEAN
	LANGUAGE plpgsql
AS $function$
DECLARE
	_ub users_base;
	_uwe users_withdraw_event;
BEGIN
	SELECT *
	FROM users_base
	WHERE tele_id = _tid
	INTO _ub;

	IF _ub.id ISNULL THEN
		RAISE EXCEPTION 'пользователь не найден';
	END IF;

	SELECT *
	FROM users_withdraw_event
	WHERE uid = _uid
	INTO _uwe;

	IF _uwe.id ISNULL THEN
		RAISE EXCEPTION 'событие по выводу не найдено';
	END IF;

	UPDATE users_base SET balance = balance + _uwe.price, conclusion = conclusion - _uwe.price WHERE _tid = _tid;
	UPDATE users_withdraw_event SET is_finished = TRUE WHERE uid = _uid;

	RETURN TRUE;
END;
$function$