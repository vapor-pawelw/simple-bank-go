-- name: CreateCustomer :one
INSERT INTO customers (first_name, last_name, email, identity_type, identity_number)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetCustomerByID :one
SELECT *
FROM customers
WHERE id = $1;

-- name: DeleteCustomer :exec
UPDATE customers
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;
