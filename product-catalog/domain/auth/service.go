package auth

import "context"

type repository interface {
	save(ctx context.Context, auth Auth) (err error)
	findByEmail(ctx context.Context, email string) (auth Auth, err error)
	// readRepository
}

// type writeRepository interface {
// 	save(ctx context.Context, item Auth) (err error)
// }

// type readRepository interface {
// 	findByEmail(ctx context.Context, email string) (item Auth, err error)
// }

type authService struct {
	repo repository
}

func newService(repo repository) authService {
	return authService{
		repo: repo,
	}
}

func (a authService) register(ctx context.Context, req Auth) (err error) {
	err = req.EncryptPassword()
	if err != nil {
		return
	}
	err = a.repo.save(ctx, req)
	if err != nil {
		return
	}

	// authItem = req
	return
}

func (a authService) login(ctx context.Context, req Auth) (item Auth, err error) {
	auth, err := a.repo.findByEmail(ctx, req.Email)
	if err != nil {
		return
	}

	ok, err := auth.ValidatePassword(req.Password)
	if err != nil {
		return req, err
	}

	if !ok {
		return req, ErrInvalidPassword
	}

	return auth, nil

}
