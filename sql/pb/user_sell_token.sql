CREATE OR REPLACE FUNCTION user_sell_token(js json, _uid character varying)
	RETURNS text
	LANGUAGE plpgsql
AS $function$
DECLARE
	_ub users_base;
	_t tokens;
	_upe users_payment_event;
BEGIN
	SELECT *
	FROM users_base
	WHERE tele_id = (js->>'tele_id')::BIGINT
	INTO _ub;

	IF _ub.id ISNULL THEN
		RAISE EXCEPTION 'пользователь не найден';
	END IF;

	SELECT *
	FROM tokens
	WHERE token_uid = js->>'token_uid'
	INTO _t;

	IF _t.id ISNULL THEN
		RAISE EXCEPTION 'токен не найден';
	END IF;

	UPDATE users_payment SET buy_nft = array_remove(buy_nft, js->>'token_uid'), sell_nft = array_append(sell_nft, js->>'token_uid') WHERE tele_id = (js->>'tele_id')::BIGINT;
	INSERT INTO users_payment_event(uid, name_token, price, tele_id) VALUES(_uid, _t.name, (js->>'token_price')::NUMERIC, (js->>'tele_id')::BIGINT) RETURNING * INTO _upe;

	RETURN _upe.uid;
END;
$function$
