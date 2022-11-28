CREATE OR REPLACE FUNCTION user_get_withdraw_event(_uid character varying)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ugwe.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT uwe.tele_id, uwe.uid, uwe.price, uwe.is_finished
			FROM users_withdraw_event uwe
			WHERE uwe.uid = _uid
		) ag
	) ugwe
	INTO _response;

	RETURN _response;
END;
$function$