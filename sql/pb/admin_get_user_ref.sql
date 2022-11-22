CREATE OR REPLACE FUNCTION admin_get_user_ref(_tid BIGINT, _tidu BIGINT)
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
			AND r.tele_id = _tidu
		)ag
	) agur
	INTO _response;

	RETURN _response;
END;
$function$
