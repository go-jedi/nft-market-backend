CREATE OR REPLACE FUNCTION user_create_referal(js json)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_r referral;
BEGIN
	SELECT *
	FROM referral
	WHERE tele_id = (js->>'tele_id')::BIGINT
	INTO _r;

	IF js->>'tele_id' = js->>'admin_referral' THEN
		RAISE EXCEPTION 'вы не можете использовать реферальную ссылку для себя';
	END IF;

	IF _r.id ISNULL THEN
		INSERT INTO referral(tele_id, tele_name, admin_referral)
		VALUES((js->>'tele_id')::BIGINT, js->>'tele_name', (js->>'admin_referral')::BIGINT);
	ELSE
		RAISE EXCEPTION 'ссылка уже была использована этим пользователем';
	END IF;
END;
$function$
