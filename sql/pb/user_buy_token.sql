CREATE OR REPLACE FUNCTION user_buy_token(js json)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_ub users_base;
BEGIN
	SELECT *
	FROM users_base
	WHERE tele_id = (js->>'tele_id')::BIGINT
	INTO _ub;

	IF _ub.id ISNULL THEN
		RAISE EXCEPTION 'пользователь не найден';
	END IF;

	UPDATE users_base SET balance = balance - (js->>'token_price')::NUMERIC WHERE tele_id = (js->>'tele_id')::BIGINT;
	UPDATE users_payment SET buy_nft = array_append(buy_nft, js->>'token_uid') WHERE tele_id = (js->>'tele_id')::BIGINT;
END;
$function$