package database

import (
	"fmt"

	"github.com/agustfricke/go-commerce/models"
)


func AddUser(user models.User) (models.User, error) {
    result, err := DB.Exec("INSERT INTO users (name, email, password, avatar) VALUES (?, ?, ?, 'default.png')", user.Name, user.Email, user.Password)
    if err != nil {
        return models.User{}, fmt.Errorf("addUser: %v", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return models.User{}, fmt.Errorf("addUser: %v", err)
    }

    insertedAlbum := models.User{
      ID: id,
      Name: user.Name,
      Email: user.Email,
      Password: user.Password,
    }

    return insertedAlbum, nil
}


