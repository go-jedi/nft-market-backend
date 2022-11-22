CREATE OR REPLACE FUNCTION user_get_min_price(_tid BIGINT)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ugmp.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT u.minim_price
			FROM users u
			WHERE u.tele_id = _tid
		) ag
	) ugmp
	INTO _response;

	RETURN _response;
END;
$function$