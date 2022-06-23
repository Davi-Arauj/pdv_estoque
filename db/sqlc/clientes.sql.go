// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: clientes.sql

package db

import (
	"context"
	"time"
)

const createCliente = `-- name: CreateCliente :one
INSERT INTO clientes
(nome,email,cpf,fone,sexo,foto,data_nascimento )
VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id, nome, email, cpf, fone, foto, sexo, data_nascimento, data_criacao, data_atualizacao
`

type CreateClienteParams struct {
	Nome           string    `json:"nome"`
	Email          string    `json:"email"`
	Cpf            string    `json:"cpf"`
	Fone           int64     `json:"fone"`
	Sexo           string    `json:"sexo"`
	Foto           string    `json:"foto"`
	DataNascimento time.Time `json:"data_nascimento"`
}

func (q *Queries) CreateCliente(ctx context.Context, arg CreateClienteParams) (Cliente, error) {
	row := q.db.QueryRowContext(ctx, createCliente,
		arg.Nome,
		arg.Email,
		arg.Cpf,
		arg.Fone,
		arg.Sexo,
		arg.Foto,
		arg.DataNascimento,
	)
	var i Cliente
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Email,
		&i.Cpf,
		&i.Fone,
		&i.Foto,
		&i.Sexo,
		&i.DataNascimento,
		&i.DataCriacao,
		&i.DataAtualizacao,
	)
	return i, err
}

const deleteCliente = `-- name: DeleteCliente :exec
DELETE FROM clientes
WHERE id=$1
`

func (q *Queries) DeleteCliente(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCliente, id)
	return err
}

const getCliente = `-- name: GetCliente :one
SELECT id, nome, email, cpf, fone, foto, sexo, data_nascimento, data_criacao, data_atualizacao FROM clientes
WHERE nome=$1 LIMIT 1
`

func (q *Queries) GetCliente(ctx context.Context, nome string) (Cliente, error) {
	row := q.db.QueryRowContext(ctx, getCliente, nome)
	var i Cliente
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Email,
		&i.Cpf,
		&i.Fone,
		&i.Foto,
		&i.Sexo,
		&i.DataNascimento,
		&i.DataCriacao,
		&i.DataAtualizacao,
	)
	return i, err
}

const listClientes = `-- name: ListClientes :many
SELECT id, nome, email, cpf, fone, foto, sexo, data_nascimento, data_criacao, data_atualizacao FROM clientes
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListClientesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListClientes(ctx context.Context, arg ListClientesParams) ([]Cliente, error) {
	rows, err := q.db.QueryContext(ctx, listClientes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Cliente{}
	for rows.Next() {
		var i Cliente
		if err := rows.Scan(
			&i.ID,
			&i.Nome,
			&i.Email,
			&i.Cpf,
			&i.Fone,
			&i.Foto,
			&i.Sexo,
			&i.DataNascimento,
			&i.DataCriacao,
			&i.DataAtualizacao,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCliente = `-- name: UpdateCliente :one
UPDATE clientes
SET  nome=$2, email=$3, cpf=$4, fone=$5, sexo=$6, foto=$7, data_nascimento=$8, data_atualizacao=now()
WHERE id=$1
RETURNING id, nome, email, cpf, fone, foto, sexo, data_nascimento, data_criacao, data_atualizacao
`

type UpdateClienteParams struct {
	ID             int64     `json:"id"`
	Nome           string    `json:"nome"`
	Email          string    `json:"email"`
	Cpf            string    `json:"cpf"`
	Fone           int64     `json:"fone"`
	Sexo           string    `json:"sexo"`
	Foto           string    `json:"foto"`
	DataNascimento time.Time `json:"data_nascimento"`
}

func (q *Queries) UpdateCliente(ctx context.Context, arg UpdateClienteParams) (Cliente, error) {
	row := q.db.QueryRowContext(ctx, updateCliente,
		arg.ID,
		arg.Nome,
		arg.Email,
		arg.Cpf,
		arg.Fone,
		arg.Sexo,
		arg.Foto,
		arg.DataNascimento,
	)
	var i Cliente
	err := row.Scan(
		&i.ID,
		&i.Nome,
		&i.Email,
		&i.Cpf,
		&i.Fone,
		&i.Foto,
		&i.Sexo,
		&i.DataNascimento,
		&i.DataCriacao,
		&i.DataAtualizacao,
	)
	return i, err
}
