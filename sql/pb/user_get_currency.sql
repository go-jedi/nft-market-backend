CREATE OR REPLACE FUNCTION user_get_currency(_tid BIGINT)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ugc.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT u.id, u.tele_id, u.currency
			FROM users u
			WHERE u.tele_id = _tid
		) ag
	) ugc
	INTO _response;

	RETURN _response;
END;
$function$