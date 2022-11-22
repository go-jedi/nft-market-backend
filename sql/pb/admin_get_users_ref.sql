CREATE OR REPLACE FUNCTION admin_get_users_ref(_tid BIGINT, _lmt INTEGER)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(agur.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT r.id, r.tele_id, r.tele_name, r.created, r.admin_referral
			FROM referral r
			WHERE r.admin_referral = _tid
			ORDER BY r.id DESC LIMIT _lmt
		) ag
	) agur
	INTO _response;

	RETURN _response;
END;
$function$