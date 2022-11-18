CREATE OR REPLACE FUNCTION user_get_lang(_tid BIGINT)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ucl.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT u.id, u.tele_id, u.lang
			FROM users u
			WHERE u.tele_id = _tid
		) ag
	) ucl
	INTO _response;

	RETURN _response;
END;
$function$