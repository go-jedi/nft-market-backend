CREATE OR REPLACE FUNCTION user_create(js json, _uid character varying)
    RETURNS BOOLEAN
    LANGUAGE plpgsql
AS $function$
DECLARE
	 _a admins;
    _u users;
BEGIN
	SELECT *
	FROM admins
	WHERE tele_id = (js->>'tele_id_admin')::BIGINT
	INTO _a;
	
	SELECT *
	FROM users
	WHERE tele_id = (js->>'tele_id')::BIGINT
	INTO _u;

	IF _a.tele_id = (js->>'tele_id')::BIGINT THEN
		INSERT INTO users(uid, tele_id, tele_name, currency, minim_price)
		VALUES(_uid, (js->>'tele_id')::BIGINT, js->>'tele_name', 'usd', 20.00);
		
		INSERT INTO users_base(tele_id)
		VALUES((js->>'tele_id')::BIGINT);

		RETURN TRUE;
	END IF;
	
	IF _u.id ISNULL THEN
		INSERT INTO users(uid, tele_id, tele_name, currency, minim_price)
		VALUES(_uid, (js->>'tele_id')::BIGINT, js->>'tele_name', 'usd', _a.minim_price);
		
		INSERT INTO users_base(tele_id)
		VALUES((js->>'tele_id')::BIGINT);

		RETURN TRUE;
	END IF;

	RETURN FALSE;
END;
$function$