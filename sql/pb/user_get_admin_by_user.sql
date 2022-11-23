CREATE OR REPLACE FUNCTION user_get_admin_by_user(_tid BIGINT)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ugabu.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT a.tele_id, u.tele_name
			FROM admins a, referral r, users u
			WHERE r.tele_id = _tid
			AND r.admin_referral = a.tele_id
			AND a.tele_id = u.tele_id
		) ag
	) ugabu
	INTO _response;

	RETURN _response;
END;
$function$