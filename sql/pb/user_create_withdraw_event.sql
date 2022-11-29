CREATE OR REPLACE FUNCTION user_create_withdraw_event(js json, _uid character varying)
	RETURNS text
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
	_uwe users_withdraw_event;
BEGIN
	SELECT *
	FROM users
	WHERE tele_id = (js->>'tele_id')::BIGINT
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь не найден';
	END IF;

	UPDATE users_base SET balance = balance - (js->>'price')::NUMERIC, conclusion = conclusion + (js->>'price')::NUMERIC WHERE tele_id = (js->>'tele_id')::BIGINT;
	INSERT INTO users_withdraw_event(tele_id, uid, price) VALUES((js->>'tele_id')::BIGINT, _uid, (js->>'price')::NUMERIC) RETURNING * INTO _uwe;
	RETURN _uwe.uid;
END;
$function$