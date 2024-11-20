package user

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/adapters/database/repository/converter"
	repoModel "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/adapters/database/repository/user/model"
	model "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/entities"
	def "github.com/NEROTEX-Team/vtb-api-2024-grpc/internal/domain/repositories"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	pool *pgxpool.Pool
	m    sync.RWMutex
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		pool: pool,
	}
}

func (r *repository) Close() {
	r.pool.Close()
}

func (r *repository) CreateUser(ctx context.Context, userData *model.CreateUserWithID) (*model.User, error) {
	r.m.Lock()
	defer r.m.Unlock()
	var user repoModel.User
	err := r.pool.QueryRow(ctx, `
		INSERT INTO users (id, first_name, last_name, email) 
		VALUES ($1, $2, $3, $4) 
		RETURNING created_at, updated_at
	`, userData.ID, userData.FirstName, userData.LastName, userData.Email).Scan(
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	user.ID = userData.ID
	user.FirstName = userData.FirstName
	user.LastName = userData.LastName
	user.Email = userData.Email

	return converter.ToUserFromRepo(&user), nil
}

func (r *repository) FetchUserById(ctx context.Context, userID string) (*model.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var user repoModel.User

	err := r.pool.QueryRow(ctx, `
		SELECT 
			id, 
			first_name, 
			last_name, 
			email, 
			created_at, 
			updated_at F
		FROM users 
		WHERE id = $1
	`, userID).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, model.ErrorUserNotFound
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repository) FetchUserList(ctx context.Context, params *model.UserListParams) (*[]model.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	rows, err := r.pool.Query(ctx, `
		SELECT 
			id,
			first_name,
			last_name,
			email,
			created_at,
			updated_at
		FROM users 
		WHERE deleted_at IS NULL 
		ORDER BY created_at
		LIMIT $1 
		OFFSET $2
	`, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []repoModel.User

	for rows.Next() {
		var user repoModel.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var modelUsers []model.User
	for _, user := range users {
		modelUsers = append(modelUsers, *converter.ToUserFromRepo(&user))
	}

	return &modelUsers, nil
}

func (r *repository) CountUsers(ctx context.Context, _ *model.UserListParams) (int64, error) {
	r.m.RLock()
	defer r.m.RUnlock()
	var count int64
	err := r.pool.QueryRow(ctx, `
		SELECT COUNT(*) 
		FROM users 
		WHERE deleted_at IS NULL
	`).Scan(&count)
	return count, err
}

func (r *repository) UpdateUserById(_ context.Context, userData *model.UpdateUser) (*model.User, error) {
	r.m.Lock()
	defer r.m.Unlock()

	var user repoModel.User
	err := r.pool.QueryRow(context.Background(), `
		UPDATE users 
		SET first_name = $2, last_name = $3, email = $4
		WHERE id = $1
		RETURNING id, first_name, last_name, email, created_at, updated_at
	`, userData.ID, userData.FirstName, userData.LastName, userData.Email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repository) DeleteUserById(ctx context.Context, userID string) error {
	r.m.Lock()
	defer r.m.Unlock()

	_, err := r.pool.Exec(ctx, `
		UPDATE users 
		SET deleted_at = TIMEZONE('utc', now()) 
		WHERE id = $1
	`, userID)
	return err
}

func (r *repository) FetchUserByEmail(ctx context.Context, email string) (*model.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	var user repoModel.User

	err := r.pool.QueryRow(ctx, `
		SELECT 
			id,
			first_name,
			last_name,
			email,
			created_at,
			updated_at 
		FROM users
		WHERE  = $1
	`, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
