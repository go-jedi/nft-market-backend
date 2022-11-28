CREATE OR REPLACE FUNCTION user_get_payment_event(_uid character varying)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ugpe.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT upe.tele_id, upe.uid, upe.name_token, upe.price
			FROM users_payment_event upe
			WHERE uid = _uid
		) ag
	) ugpe
	INTO _response;

	RETURN _response;
END;
$function$