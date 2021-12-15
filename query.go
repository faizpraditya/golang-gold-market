package main

const (
	GET_CUSTOMER_BY_ID_PREPARE = `
	SELECT
	mst_customer.id, mst_customer.first_name, mst_customer.last_name, mst_domisili.domisili
	FROM
	mst_customer JOIN mst_domisili
	ON mst_customer.domisili_id = mst_domisili.id
	WHERE mst_customer.id = :id
	ORDER BY mst_customer.id`

	UPDATE_IS_ACTIVE = `UPDATE mst_customer SET is_actived = :is_actived WHERE id = :id`
)
