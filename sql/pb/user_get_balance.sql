CREATE OR REPLACE FUNCTION user_get_balance(_tid BIGINT)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ugb.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT ub.balance
			FROM users_base ub
			WHERE ub.tele_id = _tid
		) ag
	) ugb
	INTO _response;

	RETURN _response;
END;
$function$