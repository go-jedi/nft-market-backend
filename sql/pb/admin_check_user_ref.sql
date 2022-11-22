CREATE OR REPLACE FUNCTION admin_check_user_ref(_tid BIGINT)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(acur.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT count(*)
			FROM referral r
			WHERE r.admin_referral = _tid
		) ag
	) acur
	INTO _response;

	RETURN _response;
END;
$function$