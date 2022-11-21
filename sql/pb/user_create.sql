CREATE OR REPLACE FUNCTION user_create(js json, _uid character varying)
    RETURNS void
    LANGUAGE plpgsql
AS $function$
DECLARE
    _u users;
BEGIN
    SELECT *
    FROM users
    WHERE tele_id = (js->>'tele_id')::BIGINT
    INTO _u;

    IF _u.id ISNULL THEN
        INSERT INTO users(uid, tele_id, tele_name, currency)
        VALUES(_uid, (js->>'tele_id')::BIGINT, js->>'tele_name', 'usd');

		  INSERT INTO users_base(tele_id)
		  VALUES((js->>'tele_id')::BIGINT);
    END IF;
END;
$function$