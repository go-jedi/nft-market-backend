CREATE OR REPLACE FUNCTION admin_get_user_profile(_tid BIGINT)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(agup.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT u.id, u.tele_id, u.tele_name, ub.balance, ub.is_premium, ub.verification, ub.conclusion
			FROM users u, users_base ub
			WHERE u.tele_id = _tid
			AND ub.tele_id = _tid
		) ag
	) agup
	INTO _response;

	RETURN _response;
END;
$function$