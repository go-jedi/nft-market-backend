CREATE OR REPLACE FUNCTION admin_buy_token_user(_tid BIGINT, _tknid character varying, _nprc NUMERIC, _uidpe character varying)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_ub users_base;
	_up users_payment;
BEGIN
	SELECT *
	FROM users_base
	WHERE tele_id = _tid
	INTO _ub;

	IF _ub.id ISNULL THEN
		RAISE EXCEPTION 'пользователь не найден';
	END IF;

	SELECT *
	FROM users_payment
	WHERE tele_id = _tid
	INTO _up;

	UPDATE users_payment SET sell_nft = array_remove(sell_nft, _tknid) WHERE tele_id = _tid;
	UPDATE users_base SET balance = balance + _nprc WHERE tele_id = _tid;
	UPDATE users_payment_event SET is_finished = TRUE WHERE uid = _uidpe;
END;
$function$