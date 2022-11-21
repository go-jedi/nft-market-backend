CREATE OR REPLACE FUNCTION user_get_profile(_tid BIGINT)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ugp.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT ub.id, ub.tele_id, ub.balance, ub.conclusion, ub.verification, ub.is_premium
			FROM users_base ub
			WHERE ub.tele_id = _tid
		) ag
	) ugp
	INTO _response;

	RETURN _response;
END;
$function$