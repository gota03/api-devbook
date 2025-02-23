package repositories

import (
	"api/src/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type users struct {
	db *sql.DB
}

func NewRepositoryOfUsers(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(usuario models.User) (uint8, error) {
	statement, erro := u.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")
	if erro != nil {
		log.Fatal(erro)
		return 0, nil
	}
	defer statement.Close()

	result, erro := statement.Exec(&usuario.Name, &usuario.Nick, &usuario.Email, &usuario.Password)
	if erro != nil {
		return 0, nil
	}

	lastIdInsert, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint8(lastIdInsert), nil
}

func (u users) Find(value string) ([]models.User, error) {
	value = fmt.Sprintf("%%%s%%", value)

	rows, erro := u.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?", value, value,
	)
	if erro != nil {
		return nil, erro
	}
	var users []models.User

	for rows.Next() {
		var user models.User

		erro = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		)
		if erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}
	return users, nil
}

func (u users) FindById(id uint64) (models.User, error) {
	rows, erro := u.db.Query("select id, nome, nick, email, criadoEm from usuarios where id = ?", id)
	if erro != nil {
		return models.User{}, erro
	}
	defer rows.Close()

	var user models.User

	if !rows.Next() {
		errUserNotFound := errors.New("usuário não encontrado")
		return models.User{}, errUserNotFound
	}

	erro = rows.Scan(
		&user.Id,
		&user.Name,
		&user.Nick,
		&user.Email,
		&user.CreatedAt,
	)
	if erro != nil {
		return models.User{}, erro
	}
	return user, nil
}

func (u users) Update(id uint64, user models.User) error {
	statement, erro := u.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(user.Name, user.Nick, user.Email, id)
	if erro != nil {
		return erro
	}

	return nil

}

func (u users) Delete(id uint64) error {
	statement, erro := u.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(id)
	if erro != nil {
		return erro
	}

	return nil
}

func (u users) FindByEmail(email string) (models.User, error) {
	rows, erro := u.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return models.User{}, erro
	}
	defer rows.Close()

	var user models.User

	if !rows.Next() {
		errUserNotFoundByEmail := errors.New("usuário não encontrado pelo email fornecido")
		return models.User{}, errUserNotFoundByEmail
	}

	erro = rows.Scan(&user.Id, &user.Password)
	if erro != nil {
		return models.User{}, erro
	}

	return user, nil
}

func (u users) Follow(userId, followerId uint64) error {
	statement, erro := u.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(userId, followerId)
	if erro != nil {
		return erro
	}

	return nil
}

func (u users) Unfollow(userId, followerId uint64) error {
	statement, erro := u.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(userId, followerId)
	if erro != nil {
		return erro
	}

	return nil
}

func (u users) SearchFollowers(userId uint64) ([]models.User, error) {
	rows, erro := u.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm 
		from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?
	`, userId)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		erro = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		)
		if erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil
}

func (u users) SearchFollowing(userId uint64) ([]models.User, error){
	rows, erro := u.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm 
		from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?
	`, userId)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		erro = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		)
		if erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil
}