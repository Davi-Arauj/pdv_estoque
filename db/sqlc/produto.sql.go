// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: produto.sql

package db

import (
	"context"
	"database/sql"
)

const createProduto = `-- name: CreateProduto :one
INSERT INTO produtos
(codigo_barras, nome, descricao, foto, valorpago, valorvenda, qtde, und_cod, cat_cod, scat_cod, data_atualizacao)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id, codigo_barras, nome, descricao, foto, valorpago, valorvenda, qtde, und_cod, cat_cod, scat_cod, data_criacao, data_atualizacao
`

type CreateProdutoParams struct {
	CodigoBarras    int64        `json:"codigo_barras"`
	Nome            string       `json:"nome"`
	Descricao       string       `json:"descricao"`
	Foto            string       `json:"foto"`
	Valorpago       int64        `json:"valorpago"`
	Valorvenda      int64        `json:"valorvenda"`
	Qtde            int64        `json:"qtde"`
	UndCod          int64        `json:"und_cod"`
	CatCod          int64        `json:"cat_cod"`
	ScatCod         int64        `json:"scat_cod"`
	DataAtualizacao sql.NullTime `json:"data_atualizacao"`
}

func (q *Queries) CreateProduto(ctx context.Context, arg CreateProdutoParams) (Produto, error) {
	row := q.db.QueryRowContext(ctx, createProduto,
		arg.CodigoBarras,
		arg.Nome,
		arg.Descricao,
		arg.Foto,
		arg.Valorpago,
		arg.Valorvenda,
		arg.Qtde,
		arg.UndCod,
		arg.CatCod,
		arg.ScatCod,
		arg.DataAtualizacao,
	)
	var i Produto
	err := row.Scan(
		&i.ID,
		&i.CodigoBarras,
		&i.Nome,
		&i.Descricao,
		&i.Foto,
		&i.Valorpago,
		&i.Valorvenda,
		&i.Qtde,
		&i.UndCod,
		&i.CatCod,
		&i.ScatCod,
		&i.DataCriacao,
		&i.DataAtualizacao,
	)
	return i, err
}

const deleteProduto = `-- name: DeleteProduto :exec
DELETE FROM produtos
WHERE id=$1
`

func (q *Queries) DeleteProduto(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProduto, id)
	return err
}

const getProduto = `-- name: GetProduto :one
SELECT id, codigo_barras, nome, descricao, foto, valorpago, valorvenda, qtde, und_cod, cat_cod, scat_cod, data_criacao, data_atualizacao FROM produtos
WHERE codigo_barras=$1 LIMIT 1
`

func (q *Queries) GetProduto(ctx context.Context, codigoBarras int64) (Produto, error) {
	row := q.db.QueryRowContext(ctx, getProduto, codigoBarras)
	var i Produto
	err := row.Scan(
		&i.ID,
		&i.CodigoBarras,
		&i.Nome,
		&i.Descricao,
		&i.Foto,
		&i.Valorpago,
		&i.Valorvenda,
		&i.Qtde,
		&i.UndCod,
		&i.CatCod,
		&i.ScatCod,
		&i.DataCriacao,
		&i.DataAtualizacao,
	)
	return i, err
}

const listProdutos = `-- name: ListProdutos :many
SELECT id, codigo_barras, nome, descricao, foto, valorpago, valorvenda, qtde, und_cod, cat_cod, scat_cod, data_criacao, data_atualizacao FROM produtos
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListProdutosParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProdutos(ctx context.Context, arg ListProdutosParams) ([]Produto, error) {
	rows, err := q.db.QueryContext(ctx, listProdutos, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Produto
	for rows.Next() {
		var i Produto
		if err := rows.Scan(
			&i.ID,
			&i.CodigoBarras,
			&i.Nome,
			&i.Descricao,
			&i.Foto,
			&i.Valorpago,
			&i.Valorvenda,
			&i.Qtde,
			&i.UndCod,
			&i.CatCod,
			&i.ScatCod,
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

const updateProduto = `-- name: UpdateProduto :one
UPDATE produtos
SET  nome=$2, descricao=$3, foto=$4, valorpago=$5, valorvenda=$6, qtde=$7, und_cod=$8, cat_cod=$9, scat_cod=$10, data_atualizacao=now()
WHERE codigo_barras=$1
RETURNING id, codigo_barras, nome, descricao, foto, valorpago, valorvenda, qtde, und_cod, cat_cod, scat_cod, data_criacao, data_atualizacao
`

type UpdateProdutoParams struct {
	CodigoBarras int64  `json:"codigo_barras"`
	Nome         string `json:"nome"`
	Descricao    string `json:"descricao"`
	Foto         string `json:"foto"`
	Valorpago    int64  `json:"valorpago"`
	Valorvenda   int64  `json:"valorvenda"`
	Qtde         int64  `json:"qtde"`
	UndCod       int64  `json:"und_cod"`
	CatCod       int64  `json:"cat_cod"`
	ScatCod      int64  `json:"scat_cod"`
}

func (q *Queries) UpdateProduto(ctx context.Context, arg UpdateProdutoParams) (Produto, error) {
	row := q.db.QueryRowContext(ctx, updateProduto,
		arg.CodigoBarras,
		arg.Nome,
		arg.Descricao,
		arg.Foto,
		arg.Valorpago,
		arg.Valorvenda,
		arg.Qtde,
		arg.UndCod,
		arg.CatCod,
		arg.ScatCod,
	)
	var i Produto
	err := row.Scan(
		&i.ID,
		&i.CodigoBarras,
		&i.Nome,
		&i.Descricao,
		&i.Foto,
		&i.Valorpago,
		&i.Valorvenda,
		&i.Qtde,
		&i.UndCod,
		&i.CatCod,
		&i.ScatCod,
		&i.DataCriacao,
		&i.DataAtualizacao,
	)
	return i, err
}
