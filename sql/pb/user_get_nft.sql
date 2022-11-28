CREATE OR REPLACE FUNCTION user_get_nft(_tid BIGINT)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		jsonb_build_object('nft_buy', nfb.s)
		|| jsonb_build_object('nft_sell', nfs.s)
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT t.name, t.price, t.author, t.token_uid
			FROM tokens t, users_payment up
			WHERE up.tele_id = _tid
			AND t.token_uid = ANY(up.buy_nft)
		) ag
	) nfb,
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT t.name, t.price, t.author, t.token_uid
			FROM tokens t, users_payment up
			WHERE up.tele_id = _tid
			AND t.token_uid = ANY(up.sell_nft)
		) ag
	) nfs
	INTO _response;

	RETURN _response;
END;
$function$